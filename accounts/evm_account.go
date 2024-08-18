package accounts

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/JackGod001/generate-block-chain-account/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type EvmAccount struct {
	types.BaseAccount
}

func NewEvmccount() (*EvmAccount, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	seed := bip39.NewSeed(mnemonic, "")
	privateKey, err := crypto.ToECDSA(seed[:32])
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey)

	return &EvmAccount{
		BaseAccount: types.BaseAccount{
			PublicKey:  address.Hex(),
			PrivateKey: hex.EncodeToString(crypto.FromECDSA(privateKey)),
			Mnemonic:   mnemonic,
		},
	}, nil
}
