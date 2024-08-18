package utils

import (
	"fmt"
	"github.com/JackGod001/generate-block-chain-account/factory"
	"github.com/JackGod001/generate-block-chain-account/types"
)

// BatchCreateAccounts 批量创建指定类型和数量的账户，并返回创建的账户列表
func BatchCreateAccounts(accountType types.AccountType, numAccounts int) ([]types.Account, error) {
	factory := &factory.AccountFactory{}
	accounts := make([]types.Account, 0, numAccounts)

	// 批量创建账户
	for i := 0; i < numAccounts; i++ {
		account, err := factory.CreateAccount(accountType)
		if err != nil {
			fmt.Printf("创建帐户失败 %d: %v\n", i+1, err)
			continue
		}

		accounts = append(accounts, account)
	}

	if len(accounts) == 0 {
		return nil, fmt.Errorf("无法创建任何帐户")
	}

	return accounts, nil
}
