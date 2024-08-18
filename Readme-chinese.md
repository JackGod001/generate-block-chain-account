# 生成区块链账户

[English Version](README.md)

生成区块链账户是一个 Go 语言包，提供了一种简单高效的方式来为各种区块链创建账户，包括 Solana 和兼容 EVM 的链。

## 安装

要安装此包，请使用以下命令：

```bash
go get github.com/JackGod001/generate-block-chain-account
```

## 使用方法

以下是如何使用此包的一个简单示例：

```go
package main

import (
    "fmt"
    "github.com/JackGod001/generate-block-chain-account/factory"
    "github.com/JackGod001/generate-block-chain-account/types"
)

func main() {
    factory := &factory.AccountFactory{}
    
    // 创建 Solana 账户
    solanaAccount, err := factory.CreateAccount(types.SolanaAccountType)
    if err != nil {
        fmt.Printf("创建 Solana 账户时出错：%v\n", err)
        return
    }
    fmt.Printf("Solana 公钥：%s\n", solanaAccount.GetPublicKey())

    // 创建 EVM 账户
    evmAccount, err := factory.CreateAccount(types.EvmAccountType)
    if err != nil {
        fmt.Printf("创建 EVM 账户时出错：%v\n", err)
        return
    }
    fmt.Printf("EVM 地址：%s\n", evmAccount.GetPublicKey())
}
```

## 特性

- 支持创建 Solana 和 EVM 兼容的区块链账户
- 简单易用的 API
- 高效的账户生成过程

## 支持的区块链类型

- Solana
- EVM 兼容链（如以太坊、Binance Smart Chain 等）

## 项目结构

```
generate-block-chain-account/
├── accounts/
│   ├── evm_account.go
│   └── solana_account.go
├── factory/
│   ├── account_factory.go
│   └── account_factory_test.go
├── types/
│   └── account_type.go
├── go.mod
├── README.md
└── README-chinese.md
```

## 贡献

欢迎贡献！请随时提交问题或拉取请求。

## 许可证

本项目采用 MIT 许可证。详情请参见 [LICENSE](LICENSE) 文件。
