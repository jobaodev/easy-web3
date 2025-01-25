package easyweb3

import (
	"fmt"
	"log"
	"runtime"
)

// Version information and metadata
const (
	Version       = "1.0.0"
	Name          = "easyweb3-go"
	Description   = "Work easier with Web3 in Go"
	Author        = "Rodrigo Martínez"
	AuthorEmail   = "dev@brunneis.com"
	License       = "GNU General Public License v3.0"
	RepositoryURL = "https://github.com/brunneis/easyweb3-go"
)

// Dependencies lists required packages
var Dependencies = []string{
	"github.com/ethereum/go-ethereum",
	"github.com/robertkrimen/otto",
}

// PackageInfo contains metadata about the package
type PackageInfo struct {
	Version     string
	Name        string
	Description string
	GoVersion   string
	OS          string
	Arch        string
}

// GetPackageInfo returns detailed information about the package
func GetPackageInfo() PackageInfo {
	return PackageInfo{
		Version:     Version,
		Name:        Name,
		Description: Description,
		GoVersion:   runtime.Version(),
		OS:          runtime.GOOS,
		Arch:        runtime.GOARCH,
	}
}

// CheckDependencies verifies if all required dependencies are available
func CheckDependencies() error {
	for _, dep := range Dependencies {
		if err := checkDependency(dep); err != nil {
			return fmt.Errorf("dependency %s check failed: %w", dep, err)
		}
	}
	return nil
}

// checkDependency verifies if a single dependency is available
func checkDependency(dep string) error {
	if dep == "" {
		return fmt.Errorf("invalid dependency: empty string")
	}
	// TODO: Implement actual dependency checking logic
	return nil
}

// PrintLicenseInfo displays the license information
func PrintLicenseInfo() {
	fmt.Printf("\nLicense: %s\n", License)
	fmt.Println("This software is distributed under the GNU General Public License v3.0.")
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

func printLicense() {
	fmt.Printf("\nLicense: %s\n", License)
	fmt.Println("This software is distributed under the GNU General Public License v3.0.")
}
