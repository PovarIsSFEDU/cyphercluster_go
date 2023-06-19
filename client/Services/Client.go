package Services

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/fnv"
)

type Client struct {
	Name         string
	IP           string
	hash         hash.Hash
	AddressTable map[hash.Hash]Node
	isParentNow  bool
}

func (x Client) HashedI64() int64 {
	return HashNodeI64(x.Name, x.IP)
}

func NewClient(name string, ip string, isParent bool) *Client {
	return &Client{Name: name, IP: ip, hash: HashNode256(name, ip), AddressTable: make(map[hash.Hash]Node), isParentNow: isParent}
}

func HashNode256(name string, ip string) hash.Hash {
	h := sha256.New()
	pre := []byte(name + ip)
	h.Write(pre)

	return h
}

func HashNodeI64(name string, ip string) int64 {
	h := fnv.New64a()
	pre := []byte(name + ip)
	_, err := h.Write(pre)
	if err != nil {
		return -1
	}

	return int64(h.Sum64())
}

func (x Client) AddNode(node Node) {
	hs := HashNode256(node.Name, node.IP)
	x.AddressTable[hs] = node
}

func (x Client) DeleteByAddress(name string, address string) {
	hs := HashNode256(name, address)
	delete(x.AddressTable, hs)
}

func (x Client) DeleteByHash(hs hash.Hash) {
	delete(x.AddressTable, hs)
}

func (x Client) DeleteNode(node Node) {
	delete(x.AddressTable, node.Hash)
}

func (x Client) NodeByIP(ip string) (Node, error) {
	for _, node := range x.AddressTable {
		if node.IP == ip {
			return node, nil
		}
	}
	return Node{}, fmt.Errorf("no node found for given IP: %s", ip)
}

func (x Client) NodeByName(name string) (Node, error) {
	for _, node := range x.AddressTable {
		if node.Name == name {
			return node, nil
		}
	}
	return Node{}, fmt.Errorf("no node found for given Name: %s", name)
}
