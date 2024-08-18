package accounts

import (
	"crypto/ed25519"
	"github.com/JackGod001/generate-block-chain-account/types"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip39"
)

type SolanaAccount struct {
	types.BaseAccount
}

func NewSolanaAccount() (*SolanaAccount, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	seed := bip39.NewSeed(mnemonic, "")
	privateKey := ed25519.NewKeyFromSeed(seed[:32])
	publicKey := privateKey.Public().(ed25519.PublicKey)

	// 私钥应该包括公钥部分，总共 64 字节
	fullPrivateKey := append(privateKey.Seed(), publicKey...)

	return &SolanaAccount{
		BaseAccount: types.BaseAccount{
			PublicKey:  base58.Encode(publicKey),
			PrivateKey: base58.Encode(fullPrivateKey),
			Mnemonic:   mnemonic,
		},
	}, nil
}
