package easyweb3

import (
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EasyWeb3 represents the main struct for Ethereum interactions
type EasyWeb3 struct {
	client *ethclient.Client
}

const (
	DefaultGas               = 4000000 // 4e6
	WaitLoopSeconds         = 0.1
	WaitLogLoopSeconds      = 10
	DefaultConnectionTimeout = 10
)

// InitClass initializes a new EasyWeb3 instance with an Ethereum client
func NewEasyWeb3(url string) (*EasyWeb3, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &EasyWeb3{client: client}, nil
}

// Read calls a read-only (view/pure) contract method
// Note: In Go, this would typically be implemented differently depending on the specific contract ABI
// This is a simplified version of the concept
func Read(contract interface{}, method string, parameters []interface{}) (interface{}, error) {
	// Implementation would depend on the specific contract binding
	// You would typically use abigen-generated contract bindings
	return nil, nil
}

// GetRSVFromSignature splits an Ethereum signature into its r, s, v components
func GetRSVFromSignature(signature string) (r, s string, v int64) {
	// Remove "0x" prefix if present
	if len(signature) >= 2 && signature[0:2] == "0x" {
		signature = signature[2:]
	}

	// Split signature into components
	r = signature[:64]
	s = signature[64:128]
	
	// Convert v from hex to decimal
	vHex := signature[128:]
	vBig := new(big.Int)
	vBig.SetString(vHex, 16)
	v = vBig.Int64()

	return r, s, v
}
