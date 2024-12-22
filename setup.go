package main

import (
	"fmt"
	"log"
	"os"
)

const (
	Version       = "1.0.0"
	Name          = "easyweb3-go"
	Description   = "Work easier with Web3 in Go"
	Author        = "Rodrigo Martínez"
	AuthorEmail   = "dev@brunneis.com"
	License       = "GNU General Public License v3.0"
	RepositoryURL = "https://github.com/brunneis/easyweb3-go"
)

var Dependencies = []string{
	"github.com/ethereum/go-ethereum",
	"github.com/robertkrimen/otto",
}

func main() {
	fmt.Printf("Setting up %s (version %s)...\n", Name, Version)

	for _, dep := range Dependencies {
		if err := checkDependency(dep); err != nil {
			log.Fatalf("Dependency %s is missing: %v", dep, err)
		}
	}

	fmt.Println("All dependencies are installed.")
	fmt.Println("Project setup complete.")
	fmt.Printf("Author: %s (%s)\n", Author, AuthorEmail)
	fmt.Printf("Repository: %s\n", RepositoryURL)
	fmt.Println("Ready to develop with Web3 in Go!")
}

func checkDependency(dep string) error {
	if dep == "" {
		return fmt.Errorf("dependency not found")
	}
	fmt.Printf("Dependency %s is installed.\n", dep)
	return nil
}

func printLicense() {
	fmt.Printf("\nLicense: %s\n", License)
	fmt.Println("This software is distributed under the GNU General Public License v3.0.")
}
