package cmd

import (
	"fmt"
	"strings"
	"task-tracker-cli/internal/task"
)

// svc initializes a new task service with a store for managing and persisting tasks.
var svc = task.NewService(task.NewStore())

// runTask processes task-related subcommands such as "add", "list", or "delete" for the CLI application.
func runTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: my cli task [add|list|delete]")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Please provide a task description: task-tracker-cli task add \"walk the dog\"")
			return
		}
		fmt.Println(args[1:])
		description := strings.Join(args[1:], " ")
		taskCreated, err := svc.CreateTask(description)
		if err != nil {
			fmt.Println("Failed to create task: ", err)
			return
		}
		fmt.Println("Added: ", taskCreated)
	default:
		fmt.Println("Unknown todo subcommand")
	}
}
