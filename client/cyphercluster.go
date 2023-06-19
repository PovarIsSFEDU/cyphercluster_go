package main

import (
	"cyphercluster/client/Server"
	"cyphercluster/client/Services"
)

func main() {

	name := "test2"
	ip := "127.0.0.1:8081"
	client := Services.NewClient(name, ip, false)
	s := Server.NewServer(client)
	println("Server starting...")
	println("Server started")
	s.Serve()

}
