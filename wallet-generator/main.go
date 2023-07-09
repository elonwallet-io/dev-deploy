package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	Address       string
	PrivateKeyHex string
}

func main() {
	wallet, err := generateWallet()
	if err != nil {
		panic(err)
	}

	fmt.Printf("A new wallet has been generated for you:\nAddress: %s\nPrivate key in hex format: %s", wallet.Address, wallet.PrivateKeyHex)
}

func generateWallet() (Wallet, error) {
	sk, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return Wallet{}, fmt.Errorf("failed to generate ecdsa secret key: %w", err)
	}

	pk := sk.Public()
	pkECDSA := pk.(*ecdsa.PublicKey)

	return Wallet{
		Address:       crypto.PubkeyToAddress(*pkECDSA).Hex(),
		PrivateKeyHex: hex.EncodeToString(crypto.FromECDSA(sk)),
	}, nil
}
