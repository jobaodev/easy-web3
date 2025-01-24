package main

import (
	"fmt"
	"log"
	"github.com/yourusername/easyweb3"
)

func main() {
	info := easyweb3.GetPackageInfo()
	fmt.Printf("Starting %s v%s\n", info.Name, info.Version)
	fmt.Printf("Running on %s/%s with %s\n", info.OS, info.Arch, info.GoVersion)

	if err := easyweb3.CheckDependencies(); err != nil {
		log.Fatalf("Dependency check failed: %v", err)
	}

	easyweb3.PrintLicenseInfo()
} 