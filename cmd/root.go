package cmd

import (
	"fmt"
	"os"
)

// Execute initializes and runs the CLI application, parsing the first argument as a subcommand.
// Supported subcommands include "task" and "version".
// It exits the program with an error message if no valid subcommand is provided.
func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("expected a subcommand: [task|version]")
		return
	}
	switch os.Args[1] {
	case "task":
		runTask(os.Args[2:])
	case "version":
		printVersion()
	default:
		fmt.Printf("unknown command: %s\n", os.Args[1])
	}
}
