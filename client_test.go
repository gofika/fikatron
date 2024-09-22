package fikatron

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	usdtContractAddress = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
)

func TestTRC20Balance(t *testing.T) {
	client, err := NewClient("grpc.trongrid.io:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if !assert.Nil(t, err) {
		return
	}
	defer client.Close()
	balance, err := client.TRC20Balance("TMyNZ4Bk1NfokEzc4WcEZDChVRSbdSQbZo", usdtContractAddress)
	if !assert.NoError(t, err) {
		return
	}
	balanceStr := balance.String()
	assert.True(t, balanceStr != "0")
}
