package utils

import (
	"fmt"
	"github.com/JackGod001/generate-block-chain-account/types"
	"testing"
)

func TestBatchCreateAccounts(t *testing.T) {
	tests := []struct {
		name        string
		accountType types.AccountType
		numAccounts int
		expectError bool
		expectedLen int
	}{
		{"创建5个Solana账户", types.SolanaAccountType, 5, false, 5},
		{"创建3个evm账户", types.EvmAccountType, 3, false, 3},
		{"创建0个账户", types.SolanaAccountType, 0, true, 0},
		{"创建1个账户", types.EvmAccountType, 1, false, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accounts, err := BatchCreateAccounts(tt.accountType, tt.numAccounts)

			if tt.expectError {
				if err == nil {
					t.Errorf("期望出现错误，但没有错误发生")
				}
				t.Errorf("错误信息: %v", err)
				return
			}

			if err != nil {
				t.Errorf("意外错误: %v", err)
				return
			}

			if len(accounts) != tt.expectedLen {
				t.Errorf("期望创建 %d 个账户，但实际创建了 %d 个", tt.expectedLen, len(accounts))
			}

			// 检查每个账户的属性
			for i, account := range accounts {
				fmt.Printf("账户 %d: %+v\n", i, account)
				if account.GetPublicKey() == "" {
					t.Errorf("账户 %d 的公钥为空", i)
				}
				if account.GetPrivateKey() == "" {
					t.Errorf("账户 %d 的私钥为空", i)
				}
				if account.GetMnemonic() == "" {
					t.Errorf("账户 %d 的助记词为空", i)
				}
			}
		})
	}
}
