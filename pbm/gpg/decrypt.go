package gpg

import (
	"golang.org/x/crypto/openpgp"
	"io"

	"github.com/pkg/errors"
)

func Decrypt(data io.Reader, keyring openpgp.KeyRing) (io.Reader, error) {
	message, err := openpgp.ReadMessage(data, keyring, nil, nil)

	if err != nil {
		return nil, errors.WithMessage(err, "gpg decrypt")
	}

	return message.UnverifiedBody, nil
}
