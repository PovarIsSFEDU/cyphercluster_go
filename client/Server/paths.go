package Server

import (
	"bytes"
	"crypto/rsa"
	"cyphercluster/client/Services"
	"cyphercluster/client/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
)

type rawMessage struct {
	Data      string
	TTL       int32
	Signature string
}

func (s *Server) SendMessage(w http.ResponseWriter, req *http.Request) {

	type sendRawMessage struct {
		Data      string
		TTL       int32
		Signature string
		Receiver  string
	}
	decoder := json.NewDecoder(req.Body)
	var m utils.Message
	var tmp sendRawMessage
	err := decoder.Decode(&tmp)
	if err != nil {
		println("Sending message not decoded!")
		fmt.Printf("%v \n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//res_1, er := base64.StdEncoding.DecodeString(tmp.Data)
	//if er != nil {
	//	fmt.Printf("res1 decode str 35 %v \n", err)
	//	return
	//}
	//res_2, er := base64.StdEncoding.DecodeString(tmp.Signature)
	//if er != nil {
	//	fmt.Printf("res2 decode str 40 %v \n", err)
	//	return
	//}
	m.Data = []byte(tmp.Data)
	m.TTL = tmp.TTL
	m.Signature = []byte(tmp.Signature)
	receiver, receiverErr := s.Client.NodeByName(tmp.Receiver)
	if receiverErr != nil {
		os.Exit(1)
	}
	e := s.security.EncryptMessage(&m, receiver.PublicKey)
	s.security.SignOutput(&m)
	if e != nil {
		fmt.Printf("sign 50 %v \n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		res := base64.StdEncoding.EncodeToString(m.Data)
		res2 := base64.StdEncoding.EncodeToString(m.Signature)

		tmp.Data = res
		tmp.TTL = m.TTL
		tmp.Signature = res2

		//postBody, _ := json.Marshal(m)
		postBody, _ := json.Marshal(tmp)
		responseBody := bytes.NewBuffer(postBody)
		resp, er := http.Post("http://127.0.0.1:8080/receive", "application/json", responseBody)
		if er != nil {
			fmt.Printf("An Error Occured %v \n", er)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Printf("An Error Occured %v \n", err)
			}
		}(resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
		w.Write(postBody)
	}
}

func (s *Server) ReceiveMessage(w http.ResponseWriter, req *http.Request) {
	//TODO create mechanism to take message from queue

	decoder := json.NewDecoder(req.Body)
	var t utils.Message
	var tmp rawMessage
	err := decoder.Decode(&tmp)
	if err != nil {
		println("Received message not decoded!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, er := base64.StdEncoding.DecodeString(tmp.Data)
	if er != nil {
		fmt.Printf("%v \n", err)
		return
	}
	res2, er := base64.StdEncoding.DecodeString(tmp.Signature)
	if er != nil {
		fmt.Printf("%v \n", err)
		return
	}
	t.Data = res
	t.TTL = tmp.TTL
	t.Signature = res2
	data := t.Data
	fmt.Printf("Remote address: %s\n", req.RemoteAddr)
	fmt.Printf("Remote host: %s\n", req.Host)
	//Services.NewNode("test1", "127.0.0.1:8080", s.security.PublicKey())
	requester, nodeErr := s.Client.NodeByIP(req.Host)
	if nodeErr != nil {
		fmt.Printf("%v \n", err)
	}
	isValid := s.security.VerifyInputSignature(&t, &requester)
	fmt.Printf("Received message: %s \n", data)
	fmt.Printf("Signature valid: %t \n", isValid)
	if isValid {
		decr := s.security.DecryptMessage(data)
		fmt.Printf("Decrypted message: %s \n", decr)
		w.WriteHeader(http.StatusOK)
		//	TODO - logics to resend if ttl gtz
	} else {
		println("Signatures don't match, skipping...")
		w.WriteHeader(http.StatusPreconditionFailed)
	}
}

func (s *Server) AddToCluster(w http.ResponseWriter, req *http.Request) {
	type rawNodeInfo struct {
		Name      string
		PublicKey string
	}
	decoder := json.NewDecoder(req.Body)
	var n Services.Node
	var tmp rawNodeInfo
	err := decoder.Decode(&tmp)
	if err != nil {
		println("Received message not decoded!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	n.IP = req.Host
	n.Name = tmp.Name

	bigN, ok := new(big.Int).SetString(tmp.PublicKey, 10)
	if ok != true {
		fmt.Println("Bigint decode error!")
		os.Exit(1)
	}
	n.PublicKey = &rsa.PublicKey{E: 65537, N: bigN}
	s.Client.AddNode(n)
	w.Write([]byte("Added node successfully."))
}

func (s *Server) SendAddReq(w http.ResponseWriter, _ *http.Request) {
	type rawNodeInfo struct {
		Name      string
		PublicKey string
	}
	var tmp rawNodeInfo
	tmp.Name = s.Client.Name
	tmp.PublicKey = s.security.PublicKey().N.String()

	postBody, _ := json.Marshal(tmp)
	responseBody := bytes.NewBuffer(postBody)
	resp, er := http.Post("http://127.0.0.1:8080/add", "application/json", responseBody)
	if er != nil {
		fmt.Printf("An Error Occured %v \n", er)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("An Error Occured %v \n", err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Printf("Response from adder: %s \n\n", sb)
	w.Write(body)
}

func (s *Server) GetInfo(_ http.ResponseWriter, _ *http.Request) {
	print("Hello!")
}
