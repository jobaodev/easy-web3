package easyweb3_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"

	"easyweb3"
)

func TestGetRSVFromSignature(t *testing.T) {
	sig := "0x5b2bcbbb58a123b6fdb84e3a2cd123cfcdb0e7d89f0e8a789de3ffba43a99e252bc389c7646df85e00059f2f56c539e5f41aa002da7e5298911345fc16e918d601"
	r, s, v := easyweb3.GetRSVFromSignature(sig)

	expectedR := "5b2bcbbb58a123b6fdb84e3a2cd123cfcdb0e7d89f0e8a789de3ffba43a99e25"
	expectedS := "2bc389c7646df85e00059f2f56c539e5f41aa002da7e5298911345fc16e918d6"
	expectedV := int64(1)

	assert.Equal(t, expectedR, r)
	assert.Equal(t, expectedS, s)
	assert.Equal(t, expectedV, v)
}

func TestKeccak256(t *testing.T) {
	input := "hello world"
	expectedHash := "47173285a8d7341bf90d35b2eac43aa42c6d084ea5d7077825a07c2c34be674a"

	hash := easyweb3.Keccak256(input)
	assert.Equal(t, expectedHash, hash)
}

func TestHash(t *testing.T) {
	input := "hello world"
	expectedHash := "47173285a8d7341bf90d35b2eac43aa42c6d084ea5d7077825a07c2c34be674a"

	hash := easyweb3.Hash(input)
	assert.Equal(t, expectedHash, hash)
}

func TestRecoverAddress(t *testing.T) {
	text := "hello world"
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	signature, err := crypto.Sign(accounts.TextHash([]byte(text)), privateKey)
	if err != nil {
		t.Fatalf("failed to sign: %v", err)
	}

	client := &easyweb3.EasyWeb3{}
	recoveredAddress, err := client.RecoverAddress(text, signature)
	if err != nil {
		t.Fatalf("failed to recover address: %v", err)
	}

	assert.Equal(t, address, recoveredAddress)
}
