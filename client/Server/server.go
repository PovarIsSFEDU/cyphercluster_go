package Server

import (
	"cyphercluster/client/Services"
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	anonymization *Services.Anonymization
	security      *Services.Security
	Client        *Services.Client
}

func NewServer(client *Services.Client) *Server {
	security := Services.NewSecurity()
	anonymization := Services.AnonFromHash(100, client.HashedI64())
	s := &Server{anonymization: anonymization, security: security, Client: client}
	fmt.Printf("Sec size %d\n", security.PublicKey().Size())
	//fmt.Printf("Sec E %d\n", security.PublicKey().E)
	//fmt.Printf("Sec N %d\n", security.PublicKey().N)
	println(security.PublicKey().E)
	println(security.PublicKey().N.String())
	return s
}

func (s *Server) createRouting() {
	http.HandleFunc("/send", s.SendMessage)
	http.HandleFunc("/receive", s.ReceiveMessage)
	http.HandleFunc("/info", s.GetInfo)
	http.HandleFunc("/add", s.AddToCluster)
	http.HandleFunc("/sendclreq", s.SendAddReq)
	//http.HandleFunc("/create", nil)
}

func (s *Server) Serve() {
	s.createRouting()
	err := http.ListenAndServe(":"+strings.Split(s.Client.IP, ":")[1], nil)
	if err != nil {
		fmt.Printf("Server work error! %v \n", err)
	}
}
