package Services

import (
	"crypto/sha256"
	"hash"
)

type Client struct {
	name         string
	ip           string
	hash         hash.Hash
	addressTable map[hash.Hash]Node
	isParentNow  bool
}

func initializeClient(name string, ip string, isParent bool) *Client {
	return &Client{name: name, ip: ip, hash: hashNode(name, ip), addressTable: make(map[hash.Hash]Node), isParentNow: isParent}
}

func hashNode(name string, ip string) hash.Hash {
	h := sha256.New()
	pre := []byte(name + ip)
	h.Write(pre)

	return h
}

func (x Client) addNode(node Node) {
	hs := hashNode(node.name, node.ip)
	x.addressTable[hs] = node
}

func (x Client) deleteByAddress(name string, address string) {
	hs := hashNode(name, address)
	delete(x.addressTable, hs)
}

func (x Client) deleteByHash(hs hash.Hash) {
	delete(x.addressTable, hs)
}

func (x Client) deleteNode(node Node) {
	delete(x.addressTable, node.hash)
}
