package edgefuzz

import "github.com/0xPolygon/polygon-edge/network"

func Fuzz(data []byte) int {
	network.ParseLibp2pKey(data)
	return 0
}

