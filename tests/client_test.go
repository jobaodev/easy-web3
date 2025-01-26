package client_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"your-module/client"
)

func TestNewClient(t *testing.T) {
	mockRPC := rpc.DialInProc(nil)
	mockURL := mockRPC.URL()

	config := client.NetworkConfig{
		URL:     mockURL,
		ChainID: big.NewInt(1),
		Timeout: 10,
	}

	c, err := client.NewClient(config)
	require.NoError(t, err)
	assert.NotNil(t, c)
	c.Close()
}

func TestNewClient_InvalidURL(t *testing.T) {
	config := client.NetworkConfig{
		URL:     "invalid-url",
		ChainID: big.NewInt(1),
		Timeout: 10,
	}

	_, err := client.NewClient(config)
	assert.Error(t, err)
}

func TestClient_GetBalance(t *testing.T) {
	mockRPC := rpc.DialInProc(nil)
	mockURL := mockRPC.URL()

	config := client.NetworkConfig{
		URL:     mockURL,
		ChainID: big.NewInt(1),
		Timeout: 10,
	}

	c, err := client.NewClient(config)
	require.NoError(t, err)
	defer c.Close()

	address := "0x0000000000000000000000000000000000000000"
	mockRPC.RegisterName("eth", map[string]interface{}{
		"getBalance": func(ctx context.Context, addr string, block string) (*big.Int, error) {
			return big.NewInt(1000000000000000000), nil
		},
	})

	balance, err := c.GetBalance(address)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1000000000000000000), balance)
}

func TestClient_GetBlockNumber(t *testing.T) {
	mockRPC := rpc.DialInProc(nil)
	mockURL := mockRPC.URL()

	config := client.NetworkConfig{
		URL:     mockURL,
		ChainID: big.NewInt(1),
		Timeout: 10,
	}

	c, err := client.NewClient(config)
	require.NoError(t, err)
	defer c.Close()

	mockRPC.RegisterName("eth", map[string]interface{}{
		"blockNumber": func(ctx context.Context) (string, error) {
			return "0x10", nil
		},
	})

	blockNumber, err := c.GetBlockNumber()
	require.NoError(t, err)
	assert.Equal(t, uint64(16), blockNumber)
}

func TestClient_IsContractAddress(t *testing.T) {
	mockRPC := rpc.DialInProc(nil)
	mockURL := mockRPC.URL()

	config := client.NetworkConfig{
		URL:     mockURL,
		ChainID: big.NewInt(1),
		Timeout: 10,
	}

	c, err := client.NewClient(config)
	require.NoError(t, err)
	defer c.Close()

	address := "0x0000000000000000000000000000000000000000"

	mockRPC.RegisterName("eth", map[string]interface{}{
		"getCode": func(ctx context.Context, addr string, block string) (string, error) {
			if addr == address {
				return "0x6060604052341561000f57600080fd5b", nil
			return "0x", nil
		},
	})

	isContract, err := c.IsContractAddress(address)
	require.NoError(t, err)
	assert.True(t, isContract)

	isContract, err = c.IsContractAddress("0x1111111111111111111111111111111111111111")
	require.NoError(t, err)
	assert.False(t, isContract)
}
