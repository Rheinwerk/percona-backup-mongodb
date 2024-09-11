package gpg

import (
	"golang.org/x/crypto/openpgp"
	"io"

	"github.com/pkg/errors"
)

func Encrypt(recipient *openpgp.Entity, ciphertext io.Writer) (io.WriteCloser, error) {
	plaintext, err := openpgp.Encrypt(ciphertext, []*openpgp.Entity{recipient}, nil, &openpgp.FileHints{IsBinary: true}, nil)

	if err != nil {
		return nil, errors.WithMessage(err, "gpg encrypt")
	}

	return plaintext, err
}
