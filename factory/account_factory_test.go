package factory

import (
	"fmt"
	"github.com/JackGod001/generate-block-chain-account/accounts"
	"github.com/JackGod001/generate-block-chain-account/types"
	"testing"
)

func TestAccountFactory_CreateAccount(t *testing.T) {
	factory := &AccountFactory{}

	tests := []struct {
		name        string
		accountType types.AccountType
		wantErr     bool
	}{
		{
			name:        "创建Solana账户",
			accountType: types.SolanaAccountType,
			wantErr:     false,
		},
		{
			name:        "创建evm账户",
			accountType: types.EvmAccountType,
			wantErr:     false,
		},
		{
			name:        "创建不支持的账户类型",
			accountType: types.AccountType("unsupported"),
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := factory.CreateAccount(tt.accountType)

			fmt.Printf("账户 : %+v\n", account)
			if tt.wantErr {
				if err == nil {
					t.Errorf("期望出现错误，但没有错误发生")
				}
				return
			}

			if err != nil {
				t.Errorf("意外错误: %v", err)
				return
			}

			if account == nil {
				t.Errorf("创建的账户为nil")
				return
			}

			// 检查账户类型是否正确
			switch tt.accountType {
			case types.SolanaAccountType:
				if _, ok := account.(*accounts.SolanaAccount); !ok {
					t.Errorf("期望创建SolanaAccount，但得到了不同的类型")
				}
			case types.EvmAccountType:
				if _, ok := account.(*accounts.EvmAccount); !ok {
					t.Errorf("期望创建evmAccount，但得到了不同的类型")
				}
			}

			// 检查账户的基本属性
			if account.GetPublicKey() == "" {
				t.Errorf("账户的公钥为空")
			}
			if account.GetPrivateKey() == "" {
				t.Errorf("账户的私钥为空")
			}
			if account.GetMnemonic() == "" {
				t.Errorf("账户的助记词为空")
			}
		})
	}
}
