package Services

import "hash"

type Node struct {
	name string
	ip   string
	hash hash.Hash
}
