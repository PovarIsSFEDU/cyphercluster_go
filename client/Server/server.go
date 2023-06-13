package Server

import (
	"cyphercluster/client/Services"
	"net/http"
)

type Server struct {
	anonymization *Services.Anonymization
	security      *Services.Security
	client        *Services.Client
}

func NewServer(client *Services.Client) *Server {
	security := Services.NewSecurity()
	anonymization := Services.AnonFromHash(100, client.HashedI64())
	s := &Server{anonymization: anonymization, security: security, client: client}

	return s
}

func (s *Server) createRouting() {
	http.HandleFunc("/send", s.SendMessage)
	http.HandleFunc("/receive", s.ReceiveMessage)
	http.HandleFunc("/info", s.GetInfo)
	//http.HandleFunc("/create", nil)
}

func (s *Server) Serve() {
	s.createRouting()
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		print("Server work error!")
	}
}
