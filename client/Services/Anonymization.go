package Services

import (
	"cyphercluster/client/utils"
	"math/rand"
)

type Anonymization struct {
	inputQueue        utils.Queue
	outputQueue       utils.Queue
	intervalGenerator *rand.Rand
}

func createAnonymization(size int, name string, ip string) *Anonymization {
	hashed := hashNodeI64(name, ip)
	src := rand.NewSource(hashed)
	randGen := rand.New(src)
	return &Anonymization{inputQueue: utils.Queue{Size: size}, outputQueue: utils.Queue{Size: size}, intervalGenerator: randGen}
}
