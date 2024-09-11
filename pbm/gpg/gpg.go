package gpg

import (
	"fmt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"os"
)

const KeyPath = "/etc/pbm-agent/"

const KeyFilename = "backup@centerdevice.de"
const SecretKeySuffix = ".sec"
const PublicKeySuffix = ".pub"

const SecretKeyPath = KeyPath + KeyFilename + SecretKeySuffix
const PublicKeyPath = KeyPath + KeyFilename + PublicKeySuffix

func ReadCenterDevicePublicKey() *openpgp.Entity {
	f, err := os.Open(PublicKeyPath)
	if err != nil {
		panic(fmt.Sprintf("Could not read gpg pub key: %s", PublicKeyPath))
	}
	defer f.Close()

	block, err := armor.Decode(f)
	if err != nil {
		panic(fmt.Sprintf("Could not read gpg pub key: %s", PublicKeyPath))
	}

	entity, err := openpgp.ReadEntity(packet.NewReader(block.Body))
	if err != nil {
		panic(fmt.Sprintf("Could not read gpg pub key: %s", PublicKeyPath))
	}

	return entity
}

func ReadCenterDeviceSecretKey() openpgp.EntityList {
	f, err := os.Open(SecretKeyPath)
	if err != nil {
		panic(fmt.Sprintf("Could not read gpg sec key: %s", SecretKeyPath))
	}
	defer f.Close()

	entityList, err := openpgp.ReadArmoredKeyRing(f)
	if err != nil {
		panic(fmt.Sprintf("Could not read gpg sec key: %s", SecretKeyPath))
	}

	return entityList
}
