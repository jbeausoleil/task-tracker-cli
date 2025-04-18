package cmd

import (
	"fmt"
	"sort"
	"strings"
	"task-tracker-cli/internal/task"
)

// svc is the initialized task service used by the CLI to manage and persist tasks.
var svc = task.NewService(task.NewStore())

// formatValidInput returns a formatted, sorted string representation of valid input options.
// Example output: "[done|in-progress|todo]"
func formatValidInput(options map[string]bool) string {
	var keys []string
	for key := range options {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return fmt.Sprintf("[%s]", strings.Join(keys, "|"))
}

// runTask processes task-related subcommands such as "add", "list", and "delete" from CLI arguments.
func runTask(args []string) {
	validActions := map[string]bool{
		"add":    true,
		"list":   true,
		"delete": true,
	}

	validFilters := map[string]bool{
		"done":        true,
		"todo":        true,
		"in-progress": true,
	}

	action, err := handleAction(args, validActions)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch action {
	case "add":
		handleAdd(args)
	case "list":
		handleList(args, validFilters)
	default:
		fmt.Println("Unknown task action:", action)
	}
}

// handleAction validates the provided CLI arguments against a set of valid actions.
// It returns the normalized action string if valid, or an error if invalid.
func handleAction(args []string, validActions map[string]bool) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("expected task action: %s", formatValidInput(validActions))
	}

	action := strings.ToLower(strings.TrimSpace(args[0]))

	if !validActions[action] {
		return "", fmt.Errorf("unknown action: %s, select from actions: %s", action, formatValidInput(validActions))
	}

	return action, nil
}

// handleAdd processes the "add" subcommand to create a new task.
func handleAdd(args []string) {
	if len(args) < 2 {
		fmt.Println("expected task description: task-tracker-cli task add \"walk the dog\"")
		return
	}

	description := strings.Join(args[1:], " ")
	taskCreated, err := svc.CreateTask(description)
	if err != nil {
		fmt.Println("failed to create task:", err)
		return
	}

	fmt.Println("Added:", taskCreated)
}

// handleList processes the "list" subcommand to retrieve and display tasks, optionally filtered by status.
func handleList(args []string, validFilters map[string]bool) {
	filter := ""

	if len(args) >= 2 {
		userFilter := strings.ToLower(strings.TrimSpace(args[1]))
		if !validFilters[userFilter] {
			fmt.Println("expected task list filter:", formatValidInput(validFilters))
			return
		}
		filter = userFilter
	}

	tasks := svc.ListTasks(filter)
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Printf("%-15s %-20s %-50s\n", "ID", "STATUS", "DESCRIPTION")
	fmt.Println(strings.Repeat("-", 85))
	for _, task := range tasks {
		fmt.Printf("%-15s %-20s %-50s\n", task.Id, task.Status, task.Description)
	}
}
