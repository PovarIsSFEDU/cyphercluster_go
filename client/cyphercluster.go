package main

import (
	"cyphercluster/client/Server"
	"cyphercluster/client/Services"
)

func main() {
	/*name := "test1"
	ip := "127.0.0.1:8080"
	//client := Client{name: name, ip: ip, hash: initialize(name, ip)}
	security := Services.NewSecurity()
	//security = security.GenerateKeys()
	//security.GenerateKeys()
	data := "hello"
	enc := security.EncryptMessage(data)
	msg := utils.NewMessage(enc, 120)
	security.SignOutput(msg)
	node := Services.NewNode(name, ip, security.PublicKey())
	security.VerifyInputSignature(msg, node)
	security.DecryptMessage(enc)*/

	name := "test1"
	ip := "127.0.0.1:8090"
	client := Services.NewClient(name, ip, false)
	s := Server.NewServer(client)
	println("Server starting...")
	println("Server started")
	s.Serve()

}
