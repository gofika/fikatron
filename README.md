[![codecov](https://codecov.io/gh/gofika/fikatron/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/fikatron)
[![Build Status](https://github.com/gofika/fikatron/workflows/build/badge.svg)](https://github.com/gofika/fikatron)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/fikatron)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/fikatron)](https://goreportcard.com/report/github.com/gofika/fikatron)
[![Licenses](https://img.shields.io/github/license/gofika/fikatron)](LICENSE)

# fikatron

Use GRPC to communicate with the TRON node to implement balance queries and other related operations.


## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/gofika/fikatron
```

### Example

```go
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/gofika/fikatron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	const target = "grpc.trongrid.io:50051"
	client, err := fikatron.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// calculate address
	privateKeyHex := "a3f4af5e7b28eb215f3887baa4f0e5b8d7e7c8d6e5f4c3b2a190807060504030"
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		panic(err)
	}
	_, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)
	address := fikatron.AddressFromPublicKey(publicKey.ToECDSA())

	// get account
	account, err := client.Account(address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("address: %s, balance: %d\n", address, account.Balance)

	// get TRC20-USDT balance
	const usdtContractAddress = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	balance, err := client.TRC20Balance(usdtContractAddress, usdtContractAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("address: %s, TRC20-USDT balance: %s\n", usdtContractAddress, balance.String())
}
```