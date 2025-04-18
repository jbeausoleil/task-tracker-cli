package cmd

import (
	"fmt"
	"sort"
	"strings"
	"task-tracker-cli/internal/task"
)

// svc initializes a new task service with a store for managing and persisting tasks.
var svc = task.NewService(task.NewStore())

// formatValidFilters returns a formatted, sorted string of valid task filters.
// Example output: "[done|in-progress|todo]"
func formatValidInput(filters map[string]bool) string {
	var keys []string
	for key := range filters {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return fmt.Sprintf("[%s]", strings.Join(keys, "|"))
}

// runTask processes task-related subcommands such as "add", "list", or "delete" for the CLI application.
func runTask(args []string) {
	validActions := map[string]bool{
		"add":    true,
		"list":   true,
		"delete": true,
	}

	if len(args) < 1 {
		fmt.Println("expected task action:", formatValidInput(validActions))
		return
	}

	action := strings.ToLower(strings.TrimSpace(args[0]))

	if !validActions[action] {
		fmt.Printf("unknown action: %s\nselect from actions: %s\n", action, formatValidInput(validActions))
		return
	}

	// validFilters defines the allowed task statuses for filtering.
	validFilters := map[string]bool{
		"done":        true,
		"todo":        true,
		"in-progress": true,
	}

	switch action {
	case "add":
		if len(args) < 2 {
			fmt.Println("Expected a task description: task-tracker-cli task add \"walk the dog\"")
			return
		}
		description := strings.Join(args[1:], " ")
		taskCreated, err := svc.CreateTask(description)
		if err != nil {
			fmt.Println("Failed to create task: ", err)
			return
		}
		fmt.Println("Added: ", taskCreated)
	case "list":
		filter := strings.ToLower(strings.TrimSpace(""))
		if len(args) >= 2 {
			// Validate the provided filter against validFilters.
			if !validFilters[args[1]] {
				fmt.Println("expected task list action:", formatValidInput(validFilters))
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
		fmt.Printf("%-15s %-20s %-50s\n", "ID", "STATUS", "DESCRIPTION")
		fmt.Println(strings.Repeat("-", 75))

		// Print each task
		for _, task := range tasks {
			fmt.Printf("%-15s %-20s %-50s\n", task.Id, task.Status, task.Description)
		}
	default:
		fmt.Println("Unknown todo subcommand")
	}
}
