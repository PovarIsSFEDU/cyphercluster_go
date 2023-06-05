package main

import "cyphercluster/client/Services"

func main() {
	/*name := "test1"
	ip := "127.0.0.1:8080"
	client := Client{name: name, ip: ip, hash: initialize(name, ip)}*/
	security := Services.Security{}
	security = security.GenerateKeys()
	msg := "hello"
	enc := security.EncryptMessage(msg)
	security.DecryptMessage(enc)
}
