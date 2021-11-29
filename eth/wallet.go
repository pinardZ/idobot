package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
	"log"
	"math/big"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
	client     *ethclient.Client
	chain      Chain
}

func newWallet(ctx context.Context, pkeyHex string) *wallet {
	wallet := &wallet{}

	privateKey, err := crypto.HexToECDSA(pkeyHex)
	if err != nil {
		log.Fatal(err)
	}
	wallet.privateKey = privateKey

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	wallet.address = address

	chain := &bsc{}
	wallet.chain = chain

	client, err := ethclient.DialContext(ctx, chain.GetRPCAddr())
	if err != nil {
		log.Fatal(err)
	}
	wallet.client = client

	return wallet
}

// TODO 切换网络
func switchNet(net string) {

}

// auth - 授权
func (w *wallet) auth() *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(w.privateKey, w.chain.GetChainId())
	if err != nil {
		log.Fatal(err)
	}
	return auth
}

// transfer - 转账
func (w *wallet) transfer(ctx context.Context, toAddress common.Address, value *big.Int) {
	nonce, err := w.client.PendingNonceAt(ctx, w.address)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := types.SignNewTx(w.privateKey, types.LatestSignerForChainID(w.chain.GetChainId()), &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(viper.GetInt64("gasPrice") * params.GWei),
		Gas:      viper.GetUint64("gasLimit"),
		To:       &toAddress,
		Value:    value,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = w.client.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Transaction has been sent, tx hash: %s", tx.Hash().Hex())
}
