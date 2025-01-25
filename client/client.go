package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client represents a connection to an Ethereum network
type Client struct {
	eth    *ethclient.Client
	ctx    context.Context
	cancel context.CancelFunc
}

// NetworkConfig holds the configuration for connecting to an Ethereum network
type NetworkConfig struct {
	URL      string
	ChainID  *big.Int
	Timeout  int // in seconds
	Endpoint string
}

// Common network configurations
var (
	MainnetConfig = NetworkConfig{
		URL:      "https://mainnet.infura.io/v3/YOUR-PROJECT-ID",
		ChainID:  big.NewInt(1),
		Timeout:  10,
		Endpoint: "mainnet",
	}

	GoerliConfig = NetworkConfig{
		URL:      "https://goerli.infura.io/v3/YOUR-PROJECT-ID",
		ChainID:  big.NewInt(5),
		Timeout:  10,
		Endpoint: "goerli",
	}

	SepoliaConfig = NetworkConfig{
		URL:      "https://sepolia.infura.io/v3/YOUR-PROJECT-ID",
		ChainID:  big.NewInt(11155111),
		Timeout:  10,
		Endpoint: "sepolia",
	}
)

// NewClient creates a new Ethereum client with the given configuration
func NewClient(config NetworkConfig) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())

	client, err := ethclient.Dial(config.URL)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to connect to Ethereum network: %w", err)
	}

	return &Client{
		eth:    client,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// Close terminates the client connection
func (c *Client) Close() {
	c.cancel()
	c.eth.Close()
}

// GetBalance returns the balance of the given address in Wei
func (c *Client) GetBalance(address string) (*big.Int, error) {
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid Ethereum address: %s", address)
	}

	account := common.HexToAddress(address)
	balance, err := c.eth.BalanceAt(c.ctx, account, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance, nil
}

// GetBlockNumber returns the latest block number
func (c *Client) GetBlockNumber() (uint64, error) {
	blockNumber, err := c.eth.BlockNumber(c.ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get block number: %w", err)
	}

	return blockNumber, nil
}

// IsContractAddress checks if the given address is a contract
func (c *Client) IsContractAddress(address string) (bool, error) {
	if !common.IsHexAddress(address) {
		return false, fmt.Errorf("invalid Ethereum address: %s", address)
	}

	acc := common.HexToAddress(address)
	code, err := c.eth.CodeAt(c.ctx, acc, nil)
	if err != nil {
		return false, fmt.Errorf("failed to get code at address: %w", err)
	}

	return len(code) > 0, nil
}
