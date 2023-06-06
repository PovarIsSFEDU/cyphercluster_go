package Services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io"
)

type Security struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	reader     io.Reader
}

func (x Security) PublicKey() *rsa.PublicKey {
	return x.publicKey
}

func (x Security) GenerateKeys() Security {
	x.reader = rand.Reader
	tmp, err := rsa.GenerateKey(x.reader, 1024)
	if err != nil {
		print("Error on RSA keygen!")
	}
	x.publicKey = &tmp.PublicKey
	x.privateKey = tmp
	return x
}

func (x Security) EncryptMessage(message string) string {
	fmt.Printf("Message is %s \n", message)
	encrypted, err := rsa.EncryptOAEP(sha256.New(), x.reader, x.publicKey, []byte(message), []byte("msg"))
	if err != nil {
		print("Error on RSA keygen!")

	} else {
		fmt.Printf("Encoded: %s \n", string(encrypted))
		return string(encrypted)
	}

	return ""
}

func (x Security) DecryptMessage(encrypted string) string {
	decrypted, err := rsa.DecryptOAEP(sha256.New(), x.reader, x.privateKey, []byte(encrypted), []byte("msg"))
	if err != nil {
		println("Error on RSA decrypt!")
		fmt.Printf("Error is %s \n", err.Error())
	} else {
		fmt.Printf("Decoded: %s \n", string(decrypted))
		return string(decrypted)
	}
	return ""
}
