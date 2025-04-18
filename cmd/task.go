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
		fmt.Println("expected task action: [add|list|delete]")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Expected a task description: task-tracker-cli task add \"walk the dog\"")
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
	case "list":
		filter := ""
		validFilters := map[string]bool{
			"done":        true,
			"todo":        true,
			"in-progress": true,
			"deleted":     true,
		}
		var keys []string
		for key := range validFilters {
			keys = append(keys, key)
		}
		if len(args) >= 2 {
			if !validFilters[args[1]] {
				message := fmt.Sprintf("Expected task list action [%s]", strings.Join(keys, "|"))
				fmt.Println(message)
				return
			}
			filter = args[1]
		}
		tasks := svc.ListTasks(filter)
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		// Print table header
		fmt.Printf("%-40s %-17s %-50s\n", "ID", "STATUS", "DESCRIPTION")
		fmt.Println(strings.Repeat("-", 75))

		// Print each task
		for _, task := range tasks {
			fmt.Printf("%-40s %-17s %-50s\n", task.Id, task.Status, task.Description)
		}
	default:
		fmt.Println("Unknown todo subcommand")
	}
}
