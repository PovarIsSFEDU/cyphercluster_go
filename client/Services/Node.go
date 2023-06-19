package Services

import (
	"crypto/rsa"
	"hash"
)

type Node struct {
	Name      string
	IP        string
	Hash      hash.Hash
	PublicKey *rsa.PublicKey
}

func NewNode(name string, IP string, publicKey *rsa.PublicKey) *Node {
	return &Node{Name: name, IP: IP, Hash: HashNode256(name, IP), PublicKey: publicKey}
}
