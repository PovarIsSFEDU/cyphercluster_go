package Services

import "cyphercluster/client/utils"

type Anonymization struct {
	inputQueue        utils.Queue
	outputQueue       utils.Queue
	intervalGenerator string
}
