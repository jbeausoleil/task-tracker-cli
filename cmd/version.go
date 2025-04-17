package cmd

import "fmt"

// version indicates the current version of the application in semantic versioning format.
var (
	version = "0.0.1"
	commit  = "dev"
	date    = "unknown"
)

// printVersion outputs the current application version to the console in the format "version: [version]".
func printVersion() {
	fmt.Printf("Version: %s\nCommit: %s\nBuilt at: %s\n", version, commit, date)
}
