package Services

import (
	"crypto"
	"crypto/rsa"
	"cyphercluster/client/utils"
)

func (x Security) SignOutput(m *utils.Message) {
	signature, err := rsa.SignPKCS1v15(x.reader, x.privateKey, crypto.SHA256, m.Hash())
	if err != nil {
		println("Signature error: ", err)
	} else {
		//m.Signature = string(signature)
		m.Signature = signature
		//fmt.Printf("Signed ok! Signature: %s \n", string(signature))
	}
}

func (x Security) VerifyInputSignature(m *utils.Message, sender *Node) bool {
	err := rsa.VerifyPKCS1v15(sender.PublicKey, crypto.SHA256, m.Hash(), []byte(m.Signature))
	if err != nil {
		println("Signatures don't match!: ", err)
		return false
	} else {
		//println("Signatures match! ")
		return true
	}
}

func (x Security) VerifySelf(m *utils.Message, key *rsa.PublicKey) bool {
	//err := rsa.VerifyPSS(sender.publicKey, crypto.MD5, m.Hash(), []byte(m.Signature), nil)
	err := rsa.VerifyPKCS1v15(key, crypto.SHA256, m.Hash(), []byte(m.Signature))
	if err != nil {
		//println("Signatures don't match!: ", err)
		return false
	} else {
		println("Signatures match! ")
		return true
	}
}
