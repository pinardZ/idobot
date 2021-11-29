package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
	"log"
	"math/big"
)

type Runnable interface {
	Run(ctx context.Context)
}

type dxsaleRunnable struct {
	privateKey   *ecdsa.PrivateKey
	fromAddress  common.Address
	uniAbi       abi.ABI
	ethClientMap map[string]*ethclient.Client
}

func NewDxSaleRunnable() Runnable {
	return &dxsaleRunnable{}
}

// TODO 定时任务
func (r *dxsaleRunnable) Run(ctx context.Context) {
	dxsaleContractAddress := common.HexToAddress(viper.GetString("targetContract"))
	value, _ := big.NewFloat(viper.GetFloat64("buyingBnbOrEthAmount") * params.Ether).Int(nil)
	// TODO 估算gas
	gas := viper.GetUint64("estimateGas")

	if viper.GetUint64("gasLimit") < gas {
		log.Println("config gas limit less than estimate gas ", gas, "auto set to estimate gasLimit")
		viper.Set("gasLimit", gas)
	}

	log.Println("EstimateGas", gas, "ready to transfer")

	privateKey := viper.GetString("privateKey")
	wallet := newWallet(ctx, privateKey)
	wallet.transfer(ctx, dxsaleContractAddress, value)
}
