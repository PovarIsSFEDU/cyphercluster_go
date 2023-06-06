package utils

import (
	"crypto/sha256"
)

type Message struct {
	Data      string
	TTL       int32
	Signature string
}

func NewMessage(data string, TTL int32) *Message {
	return &Message{Data: data, TTL: TTL}
}

func (m *Message) Hash() []byte {
	h := sha256.New()
	pre := []byte(m.Data + string(m.TTL))
	h.Write(pre)
	return h.Sum(nil)
}
