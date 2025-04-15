package cmd

import "fmt"

// version shows CLI
const version = "0.1.0"

func printVersion() {
	fmt.Printf("version: %s\n", version)
}
