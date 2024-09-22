package fikatron

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/assert"
)

func TestAddress(t *testing.T) {
	privateKeyHex := "a3f4af5e7b28eb215f3887baa4f0e5b8d7e7c8d6e5f4c3b2a190807060504030"
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		panic(err)
	}
	_, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)
	address := AddressFromPublicKey(publicKey.ToECDSA())
	assert.Equal(t, "TDbka3asbgtMfYnM7HVkmdjHT7brWGZHX6", address)
}
