package encrypt

import (
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"

	"github.com/percona/percona-backup-mongodb/pbm/errors"
)

const (
	KeyPath         = "/etc/pbm-agent/"
	KeyFilename     = "backup@centerdevice.de"
	SecretKeySuffix = ".sec"
	PublicKeySuffix = ".pub"
	SecretKeyPath   = KeyPath + KeyFilename + SecretKeySuffix
	PublicKeyPath   = KeyPath + KeyFilename + PublicKeySuffix
)

func ignore(name string) bool {
	return false ||
		strings.HasSuffix(name, "pbm.init") || // pbm initialization in the s3 root
		strings.HasPrefix(name, ".pbm.restore") || // restore: coordination of nodes
		strings.HasSuffix(name, "metadata.json") || // restore: metadata
		strings.HasSuffix(name, ".pbm.json") || // backup: metadata
		strings.HasSuffix(name, "filelist.pbm") // backup: list of files
}

func EncryptFileWithGPG(name string, data io.Reader) (io.Reader, error) {
	if ignore(name) {
		return data, nil
	}

	return encryptWithGPG(data)
}

func encryptWithGPG(data io.Reader) (io.Reader, error) {
	publicKey, err := readCenterDevicePublicKey()
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()

	go func() {
		defer pw.Close()

		encryptedWriter, err := openpgp.Encrypt(pw, []*openpgp.Entity{publicKey}, nil, nil, nil)
		if err != nil {
			_ = pw.CloseWithError(err)
			return
		}
		defer encryptedWriter.Close()

		_, err = io.Copy(encryptedWriter, data)
		if err != nil {
			pw.CloseWithError(err)
		}
	}()
	return pr, nil
}

func readCenterDevicePublicKey() (*openpgp.Entity, error) {
	f, err := os.Open(PublicKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "open public key")
	}
	defer f.Close()

	block, err := armor.Decode(f)
	if err != nil {
		return nil, errors.Wrap(err, "decode public key")
	}

	entity, err := openpgp.ReadEntity(packet.NewReader(block.Body))
	if err != nil {
		return nil, errors.Wrap(err, "read public key")
	}

	return entity, nil
}

func DecryptFileWithGPG(name string, data io.ReadCloser) (io.ReadCloser, error) {
	if ignore(name) {
		return data, nil
	}

	return decryptWithGPG(data)
}

func decryptWithGPG(encryptedData io.Reader) (io.ReadCloser, error) {
	privateKey, err := readCenterDevicePrivateKey()
	if err != nil {
		return nil, err
	}

	md, err := openpgp.ReadMessage(encryptedData, privateKey, nil, nil)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(md.UnverifiedBody), nil
}

func readCenterDevicePrivateKey() (openpgp.EntityList, error) {
	f, err := os.Open(SecretKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "could not read gpg sec key")
	}
	defer f.Close()

	entityList, err := openpgp.ReadArmoredKeyRing(f)
	if err != nil {
		return nil, errors.Wrap(err, "could not read gpg sec key")
	}
	return entityList, nil
}
