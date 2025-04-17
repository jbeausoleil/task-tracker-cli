package cmd

import "fmt"

// version indicates the current version of the application in semantic versioning format.
const version = "0.1.0"

// printVersion outputs the current application version to the console in the format "version: [version]".
func printVersion() {
	fmt.Printf("version: %s\n", version)
}
