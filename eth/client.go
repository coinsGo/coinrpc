package eth

import (
	"github.com/go-develop/coinrpc/rpc"
)

type EthClient struct {
	ConnCfg *rpc.ConnConfig
}

func NewEthClient(connCfg *rpc.ConnConfig) *EthClient {
	return &EthClient{connCfg}
}
