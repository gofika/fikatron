package fikatron

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"

	"github.com/gofika/fikatron/api"
	"github.com/gofika/fikatron/core"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Client client for the TRON RPC
type Client struct {
	cli       *grpc.ClientConn
	walletCli api.WalletClient
}

// NewClient creates a new client
func NewClient(target string, opts ...grpc.DialOption) (*Client, error) {
	cli, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}
	walletCli := api.NewWalletClient(cli)
	return &Client{cli: cli, walletCli: walletCli}, nil
}

// Close closes the client
func (c *Client) Close() error {
	return c.cli.Close()
}

// Target returns the target string of the ClientConn.
func (c *Client) Target() string {
	return c.cli.Target()
}

// WalletClient returns the wallet client
// If you need to call other RPCs, you can use this to get the client
func (c *Client) WalletClient() api.WalletClient {
	return c.walletCli
}

// TRC20Call calls a TRC20 contract
func (c *Client) TRC20Call(from, contractAddress string, data []byte, constant bool, feeLimit uint64) (*api.TransactionExtention, error) {
	var err error
	ownerAddress := []byte{0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if len(from) > 0 {
		ownerAddress, err = fromBase58(from)
		if err != nil {
			return nil, err
		}
	}
	contractAddressBytes, err := fromBase58(contractAddress)
	if err != nil {
		return nil, err
	}
	req := &core.TriggerSmartContract{
		OwnerAddress:    ownerAddress,
		ContractAddress: contractAddressBytes,
		Data:            data,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var resp *api.TransactionExtention
	if constant {
		resp, err = c.walletCli.TriggerConstantContract(ctx, req)
	} else {
		resp, err = c.walletCli.TriggerContract(ctx, req)
	}
	if err != nil {
		return nil, err
	}
	if resp.Result.Code != 0 {
		return nil, fmt.Errorf(string(resp.Result.Message))
	}
	if !constant && feeLimit > 0 {
		resp.Transaction.RawData.FeeLimit = int64(feeLimit)
		raw, err := proto.Marshal(resp.Transaction.GetRawData())
		if err != nil {
			return nil, err
		}
		hash := sha256.Sum256(raw)
		resp.Txid = hash[:]
	}

	return resp, nil
}

// TRC20Balance gets the balance of a TRC20 token
func (c *Client) TRC20Balance(address, contractAddress string) (*big.Int, error) {
	addr, err := fromBase58(address)
	if err != nil {
		return nil, err
	}
	methodID := []byte{0x70, 0xa0, 0x82, 0x31} // balanceOf(address) method keccak256 hash first 4 bytes
	paddedAddress := make([]byte, 32)
	copy(paddedAddress[11:], addr)
	data := append(methodID, paddedAddress...)
	resp, err := c.TRC20Call("", contractAddress, data, true, 0)
	if err != nil {
		return nil, err
	}
	results := resp.GetConstantResult()
	if len(results) == 0 {
		return nil, fmt.Errorf("no result from TRC20Call")
	}
	res := results[0]
	return new(big.Int).SetBytes(res), nil
}

// Account gets the account information
func (c *Client) Account(address string) (*core.Account, error) {
	addr, err := fromBase58(address)
	if err != nil {
		return nil, err
	}
	req := &core.Account{
		Address: addr,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return c.walletCli.GetAccount(ctx, req)
}
