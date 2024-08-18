# Generate Block Chain Account
[中文文档](Readme-chinese.md)
Generate Block Chain Account is a Go package that provides a simple and efficient way to create blockchain accounts for various chains, including Solana and EVM-compatible chains.

## Installation

To install this package, use the following command:

```bash
go get github.com/JackGod001/generate-block-chain-account
```

## Usage

Here's a quick example of how to use this package:

```go
package main

import (
    "fmt"
    "github.com/JackGod001/generate-block-chain-account/factory"
    "github.com/JackGod001/generate-block-chain-account/types"
)

func main() {
    factory := &factory.AccountFactory{}
    
    // Create a Solana account
    solanaAccount, err := factory.CreateAccount(types.SolanaAccountType)
    if err != nil {
        fmt.Printf("Error creating Solana account: %v\n", err)
        return
    }
    fmt.Printf("Solana Public Key: %s\n", solanaAccount.GetPublicKey())

    // Create an Evm account
    ethAccount, err := factory.CreateAccount(types.EvmAccountType)
    if err != nil {
        fmt.Printf("Error creating Evm account: %v\n", err)
        return
    }
    fmt.Printf("Evm Address: %s\n", ethAccount.GetPublicKey())
}
```
## Features

- Support for creating Solana and EVM-compatible blockchain accounts
- Simple and easy-to-use API
- Efficient account generation process

## Supported Blockchain Types

- Solana
- EVM-compatible chains (such as Ethereum, Binance Smart Chain, etc.)

## Project Structure

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

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
