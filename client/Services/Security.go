package Services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"cyphercluster/client/utils"
	"io"
)

type Security struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	reader     io.Reader
}

func NewSecurity() *Security {
	x := &Security{}
	x.reader = rand.Reader
	tmp, err := rsa.GenerateKey(x.reader, 1024)
	if err != nil {
		//print("Error on RSA keygen!")
	}
	x.publicKey = &tmp.PublicKey
	x.privateKey = tmp
	return x
}

func (x Security) PublicKey() *rsa.PublicKey {
	return x.publicKey
}

func (x Security) EncryptMessage(message *utils.Message, receiverPK *rsa.PublicKey) error {
	//fmt.Printf("Message is %s \n", message.Data)
	encrypted, err := rsa.EncryptOAEP(sha256.New(), x.reader, receiverPK, message.Data, []byte("msg"))
	if err != nil {
		print("Error on RSA keygen!")
		return err
	} else {
		//fmt.Printf("Encoded: %s \n", string(encrypted))
		message.Data = encrypted
		return nil
	}
}

func (x Security) DecryptMessage(encrypted []byte) string {

	decrypted, err := rsa.DecryptOAEP(sha256.New(), x.reader, x.privateKey, encrypted, []byte("msg"))
	if err != nil {
		println("Error on RSA decrypt!")
		//fmt.Printf("Error is %s \n", err.Error())
	} else {
		//fmt.Printf("Decoded: %s \n", string(decrypted))
		return string(decrypted)
	}
	return ""
}
