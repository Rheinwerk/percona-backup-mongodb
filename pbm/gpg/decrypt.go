package gpg

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/openpgp"
	"io"
)

func Decrypt(data io.Reader, keyring openpgp.KeyRing) (io.Reader, error) {
	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()
		_, err := io.Copy(pw, data)
		if err != nil {
			pw.CloseWithError(err)
		}
	}()

	done := make(chan struct{})
	var message *openpgp.MessageDetails
	var err error

	go func() {
		message, err = openpgp.ReadMessage(pr, keyring, nil, nil)
		close(done)
	}()

	select {
	case <-done:
		if err != nil {
			return nil, errors.New("gpg decrypt failed: " + err.Error())
		}
		return message.UnverifiedBody, nil
	}
}
