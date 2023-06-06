package main

import (
	"cyphercluster/client/Services"
	"cyphercluster/client/utils"
)

func main() {
	name := "test1"
	ip := "127.0.0.1:8080"
	//client := Client{name: name, ip: ip, hash: initialize(name, ip)}
	security := Services.Security{}
	security = security.GenerateKeys()
	data := "hello"
	enc := security.EncryptMessage(data)
	msg := utils.NewMessage(enc, 120)
	security.SignOutput(msg)
	node := Services.NewNode(name, ip, security.PublicKey())
	security.VerifyInputSignature(msg, node)
	security.DecryptMessage(enc)
}
