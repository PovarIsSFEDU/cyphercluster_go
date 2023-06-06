package Services

import (
	"crypto"
	"crypto/rsa"
	"cyphercluster/client/utils"
	"fmt"
)

func (x Security) SignOutput(m *utils.Message) {
	signature, err := rsa.SignPSS(x.reader, x.privateKey, crypto.SHA256, m.Hash(), nil)
	//signature, err := x.privateKey.Sign(x.reader, []byte((m.Data)), nil)
	if err != nil {
		println("Signature error: ", err)
	} else {
		m.Signature = string(signature)
		fmt.Printf("Signed ok! Signature: %s \n", string(signature))
	}
}

func (x Security) VerifyInputSignature(m *utils.Message, sender *Node) {
	err := rsa.VerifyPSS(sender.publicKey, crypto.SHA256, m.Hash(), []byte(m.Signature), nil)
	if err != nil {
		println("Signatures don't match!: ", err)
	} else {
		println("Signatures match! ")
	}
}
