package Services

import (
	"crypto/sha256"
	"hash"
	"hash/fnv"
)

type Client struct {
	name         string
	ip           string
	hash         hash.Hash
	addressTable map[hash.Hash]Node
	isParentNow  bool
}

func initializeClient(name string, ip string, isParent bool) *Client {
	return &Client{name: name, ip: ip, hash: HashNode256(name, ip), addressTable: make(map[hash.Hash]Node), isParentNow: isParent}
}

func HashNode256(name string, ip string) hash.Hash {
	h := sha256.New()
	pre := []byte(name + ip)
	h.Write(pre)

	return h
}

func hashNodeI64(name string, ip string) int64 {
	h := fnv.New64a()
	pre := []byte(name + ip)
	_, err := h.Write(pre)
	if err != nil {
		return -1
	}

	return int64(h.Sum64())
}

func (x Client) addNode(node Node) {
	hs := HashNode256(node.Name, node.IP)
	x.addressTable[hs] = node
}

func (x Client) deleteByAddress(name string, address string) {
	hs := HashNode256(name, address)
	delete(x.addressTable, hs)
}

func (x Client) deleteByHash(hs hash.Hash) {
	delete(x.addressTable, hs)
}

func (x Client) deleteNode(node Node) {
	delete(x.addressTable, node.Hash)
}
