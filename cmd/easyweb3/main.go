package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jobaodev/easyweb3"
	"github.com/jobaodev/easyweb3/client"
)

func main() {
	info := easyweb3.GetPackageInfo()
	fmt.Printf("Starting %s v%s\n", info.Name, info.Version)
	fmt.Printf("Running on %s/%s with %s\n", info.OS, info.Arch, info.GoVersion)

	if err := easyweb3.CheckDependencies(); err != nil {
		log.Fatalf("Dependency check failed: %v", err)
	}

	// Initialize Ethereum client with Sepolia testnet
	config := client.SepoliaConfig
	config.URL = os.Getenv("ETH_NODE_URL") // Get URL from environment variable

	eth, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Ethereum client: %v", err)
	}
	defer eth.Close()

	// Get latest block number
	blockNum, err := eth.GetBlockNumber()
	if err != nil {
		log.Fatalf("Failed to get block number: %v", err)
	}
	fmt.Printf("Current block number: %d\n", blockNum)

	// Example: Check USDC contract on Sepolia
	usdcAddress := "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238" // Sepolia USDC
	isContract, err := eth.IsContractAddress(usdcAddress)
	if err != nil {
		log.Fatalf("Failed to check contract: %v", err)
	}
	fmt.Printf("Is USDC a contract? %v\n", isContract)

	easyweb3.PrintLicenseInfo()
}
