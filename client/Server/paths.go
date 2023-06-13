package Server

import (
	"bytes"
	"cyphercluster/client/Services"
	"cyphercluster/client/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) SendMessage(w http.ResponseWriter, req *http.Request) {
	type rawMessage struct {
		Data      string
		TTL       int32
		Signature string
	}

	decoder := json.NewDecoder(req.Body)
	var m utils.Message
	var tmp rawMessage
	err := decoder.Decode(&tmp)
	if err != nil {
		println("Received message not decoded 1!")
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

	e := s.security.EncryptMessage(&m)
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
		resp, er := http.Post("http://127.0.0.1:8090/receive", "application/json", responseBody)
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
	type rawMessage struct {
		Data      string
		TTL       int32
		Signature string
	}
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
	isValid := s.security.VerifyInputSignature(&t, Services.NewNode("test1", "127.0.0.1:8090", s.security.PublicKey()))
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

func (s *Server) GetInfo(_ http.ResponseWriter, _ *http.Request) {
	print("Hello!")
}
