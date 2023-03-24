package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	Client  *ethclient.Client
	ChainID *big.Int
	Signer  types.Signer
}

func NewChain(url string) (*Chain, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Failed to get chain ID: %v", err)
		return nil, err
	}

	signer := types.NewEIP155Signer(chainID)

	return &Chain{
		Client:  client,
		ChainID: chainID,
		Signer:  signer,
	}, nil
}

type Account struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

func NewAccount(initialPrv string) (*Account, error) {
	privateKey, err := crypto.HexToECDSA(initialPrv)
	if err != nil {
		log.Printf("Failed to get private key: %v", err)
		return nil, err
	}

	// get address from private key
	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &Account{
		privateKey: privateKey,
		address:    address,
	}, nil
}

type Executor struct {
	account *Account
	chain   *Chain
}

// NewExecutor creates a new Executor instance
func NewExecutor(url string, initialPrv string) (*Executor, error) {
	chain, err := NewChain(url)
	if err != nil {
		return nil, err
	}

	account, err := NewAccount(initialPrv)
	if err != nil {
		return nil, err
	}

	return &Executor{
		account: account,
		chain:   chain,
	}, nil
}

// Client returns the underlying ethclient.Client
func (ex *Executor) Client() *ethclient.Client {
	return ex.chain.Client
}

// Get nonce of initial address
func (ex *Executor) getNonce() (uint64, error) {
	client := ex.Client()
	nonce, err := client.PendingNonceAt(context.Background(), ex.account.address)
	if err != nil {
		log.Printf("Failed to get nonce: %v", err)
		return 0, err
	}
	return nonce, nil
}

// Get suggested gas price
func (ex *Executor) getSuggestedGasPrice() (*big.Int, error) {
	client := ex.Client()
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Failed to get gas price: %v", err)
		return nil, err
	}
	return gasPrice, nil
}

// Get constant gas price
func (ex *Executor) getGasPrice() *big.Int {
	return big.NewInt(100000001 * 2)
}

// Transfer amount ETH to address
func (ex *Executor) transfer2Address(address string, amount int64) (string, error) {

	nonce, err := ex.getNonce()
	if err != nil {
		return "", err
	}
	destinationAddress := common.HexToAddress(address)

	tx := types.NewTransaction(nonce, destinationAddress, big.NewInt(amount), 210000, ex.getGasPrice(), nil)

	signedTx, err := types.SignTx(tx, ex.chain.Signer, ex.account.privateKey)
	if err != nil {
		log.Printf("Failed to sign transaction: %v", err)
		return "", err
	}

	client := ex.Client()
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("Failed to send transaction: %v", err)
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
