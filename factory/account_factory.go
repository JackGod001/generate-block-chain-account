package factory

import (
	"fmt"
	"github.com/JackGod001/generate-block-chain-account/accounts"
	"github.com/JackGod001/generate-block-chain-account/types"
)

type AccountFactory struct{}

func (f *AccountFactory) CreateAccount(accountType types.AccountType) (types.Account, error) {
	switch accountType {
	case types.SolanaAccountType:
		return accounts.NewSolanaAccount()
	case types.EvmAccountType:
		return accounts.NewEvmccount()
	default:
		return nil, fmt.Errorf("unsupported account type: %s", accountType)
	}
}
