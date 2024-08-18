package types

type AccountType string

const (
	EvmAccountType    AccountType = "evm"
	SolanaAccountType AccountType = "solana"
)

type Account interface {
	GetPublicKey() string
	GetPrivateKey() string
	GetMnemonic() string
}

type BaseAccount struct {
	PublicKey  string
	PrivateKey string
	Mnemonic   string
}

func (a *BaseAccount) GetPublicKey() string {
	return a.PublicKey
}

func (a *BaseAccount) GetPrivateKey() string {
	return a.PrivateKey
}

func (a *BaseAccount) GetMnemonic() string {
	return a.Mnemonic
}
