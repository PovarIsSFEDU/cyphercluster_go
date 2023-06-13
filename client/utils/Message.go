package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

type Message struct {
	Data      []byte
	TTL       int32
	Signature []byte
}

func NewMessage(data string, TTL int32) *Message {
	res, er := base64.StdEncoding.DecodeString(data)
	if er != nil {
		return nil
	}
	return &Message{Data: res, TTL: TTL}
}

func (m *Message) Hash() []byte {
	hash := sha256.Sum256([]byte(string(m.Data) + string(m.TTL)))
	return hash[:]
}
