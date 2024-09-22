package fikatron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"math/big"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	TronBytePrefix = 0x41
	AddressLength  = 20
)

var (
	ErrInvalidAddress = errors.New("invalid address")
)

// AddressFromPublicKey convert public key to Tron address
func AddressFromPublicKey(publicKey *ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(*publicKey)

	addressTron := make([]byte, 0)
	addressTron = append(addressTron, TronBytePrefix)
	addressTron = append(addressTron, address.Bytes()...)

	if addressTron[0] == 0 {
		return new(big.Int).SetBytes(addressTron).String()
	}

	h := doubleSHA256(addressTron)

	inputCheck := addressTron
	inputCheck = append(inputCheck, h[:4]...)

	return base58.Encode(inputCheck)
}

func fromBase58(address string) ([]byte, error) {
	decoded, version, err := base58.CheckDecode(address)
	if err != nil {
		return nil, err
	}
	if version != TronBytePrefix {
		return nil, ErrInvalidAddress
	}
	if len(decoded) != AddressLength {
		return nil, ErrInvalidAddress
	}
	return append([]byte{TronBytePrefix}, decoded...), nil
}

func doubleSHA256(input []byte) []byte {
	h := sha256.Sum256(input)
	h = sha256.Sum256(h[:])
	return h[:]
}
