package main

import (
	"claimer/dist"
	"context"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"claimer/token"
)

type Claimer struct {
	Executor
	distContract  *dist.Dist
	tokenContract *token.Token
}

// Builder distributor contract
// https://arbiscan.io/address/0x67a24ce4321ab3af51c2d0a4801c3e111d88c9d9
func (cl *Claimer) buildDistributor() error {
	client := cl.Client()
	distContract, err := dist.NewDist(common.HexToAddress("0x67a24CE4321aB3aF51c2D0a4801c3E111D88C9d9"), client)
	if err != nil {
		log.Printf("Failed to build distributor contract: %v", err)
		return err
	}
	cl.distContract = distContract
	return nil
}

// Builder token contract
// https://arbiscan.io/address/0x912ce59144191c1204e64559fe8253a0e49e6548
func (cl *Claimer) buildToken() error {
	client := cl.Client()
	tokenContract, err := token.NewToken(common.HexToAddress("0x912CE59144191C1204E64559FE8253a0e49E6548"), client)
	if err != nil {
		log.Printf("Failed to build token contract: %v", err)
		return err
	}
	cl.tokenContract = tokenContract
	return nil
}

func (cl *Claimer) claim() (string, error) {

	auth := bind.NewKeyedTransactor(cl.account.privateKey)
	nonce, err := cl.getNonce()
	if err != nil {
		log.Printf("Failed to get nonce: %v", err)
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(600000) // in units
	auth.GasPrice = cl.getGasPrice()

	tx, err := cl.distContract.Claim(auth)
	if err != nil {
		log.Printf("Failed to claim: %v", err)
		return "", err
	}

	signedTx, err := types.SignTx(tx, cl.chain.Signer, cl.account.privateKey)
	if err != nil {
		log.Printf("Failed to sign transaction: %v", err)
		return "", err
	}

	client := cl.Client()
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("Failed to send transaction: %v", err)
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (cl *Claimer) withdrawTokens(to string, amount float64) (string, error) {
	auth := bind.NewKeyedTransactor(cl.account.privateKey)
	nonce, err := cl.getNonce()
	if err != nil {
		log.Printf("Failed to get nonce: %v", err)
		return "", err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(600000) // in units
	auth.GasPrice = cl.getGasPrice()

	decimals, err := cl.tokenContract.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get decimals: %v", err)
		return "", err
	}

	toAddress := common.HexToAddress(to)
	amountBigInt := big.NewInt(int64(amount * math.Pow(10, float64(decimals))))
	tx, err := cl.tokenContract.Transfer(auth, toAddress, amountBigInt)
	if err != nil {
		log.Printf("Failed to transfer tokens: %v", err)
		return "", err
	}

	signedTx, err := types.SignTx(tx, cl.chain.Signer, cl.account.privateKey)
	if err != nil {
		log.Printf("Failed to sign transaction: %v", err)
		return "", err
	}

	client := cl.Client()
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("Failed to send transaction: %v", err)
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
