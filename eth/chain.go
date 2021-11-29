package eth

import "math/big"

type Chain interface {
	GetRPCAddr() string
	GetChainType() string
	GetChainId() *big.Int
}

var _ Chain = (*bsc)(nil)

type bsc struct{}

func (b bsc) GetRPCAddr() string {
	return BSCRPCAddr
}

func (b bsc) GetChainType() string {
	return BSCChainType
}

func (b bsc) GetChainId() *big.Int {
	return big.NewInt(BSCChainId)
}
