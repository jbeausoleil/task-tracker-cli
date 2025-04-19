package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"task-tracker-cli/internal/task"
)

// svc is the initialized task service used by the CLI to manage and persist tasks.
var svc = task.NewService(task.NewStore())

// formatValidInput returns a formatted, sorted string representation of valid input options.
// Example output: "[done|in-progress|todo]".
func formatValidInput(options map[string]struct{}) string {
	var keys []string
	for key := range options {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return fmt.Sprintf("[%s]", strings.Join(keys, "|"))
}

// runTask processes task-related subcommands such as "add", "list", and "delete" from CLI arguments.
func runTask(args []string) {
	if len(args) < 1 {
		fmt.Println("expected task action")
		return
	}

	validActions := map[string]struct{}{
		"add":    {},
		"list":   {},
		"delete": {},
		"update": {},
	}

	validFilters := map[string]struct{}{
		"done":        {},
		"todo":        {},
		"in-progress": {},
	}

	action := normalizeAction(args[0])

	// Handle dynamic "mark-" commands first
	if strings.HasPrefix(action, "mark-") {
		handleMarkStatus(args)
		return
	}

	// Validate fixed actions
	if _, exists := validActions[action]; !exists {
		fmt.Printf("unknown action: %s, select from actions: %s\n", action, formatValidInput(validActions))
		return
	}

	handlers := map[string]func(){
		"add":  func() { handleAdd(args) },
		"list": func() { handleList(args, validFilters) },
		"delete": func() {
			result, err := handleDelete(args)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		},
		"update": func() {
			handleUpdateDescription(args)
		},
	}

	handler := handlers[action]
	handler()
}

// normalizeAction trims and lowercases a CLI action string.
func normalizeAction(input string) string {
	return strings.ToLower(strings.TrimSpace(input))
}

// handleAction validates the provided CLI arguments against a set of valid actions.
// It returns the normalized action string if valid, or an error if invalid.
func handleAction(args []string, validActions map[string]struct{}) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("expected task action: %s", formatValidInput(validActions))
	}

	action := strings.ToLower(strings.TrimSpace(args[0]))

	// Check for fixed valid actions
	if _, exists := validActions[action]; exists {
		return action, nil
	}

	// Special dynamic case: allow "mark-<status>"
	if strings.HasPrefix(action, "mark-") {
		return action, nil
	}

	// Otherwise invalid
	return "", fmt.Errorf("unknown action: %s, select from actions: %s", action, formatValidInput(validActions))
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

	fmt.Println("Added task:", taskCreated)
}

// handleList processes the "list" subcommand to retrieve and display tasks, optionally filtered by status.
func handleList(args []string, validFilters map[string]struct{}) {
	filter := ""

	if len(args) >= 2 {
		userFilter := strings.ToLower(strings.TrimSpace(args[1]))
		if _, exists := validFilters[userFilter]; !exists {
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

// handleDelete processes the "delete" subcommand to delete a task by its ID.
//
// It expects the task ID to be provided as the second argument in the args slice.
// The function prompts the user for confirmation before deleting the task.
// If the user confirms by typing "y" or "yes" (case-insensitive), it proceeds with deletion.
// Otherwise, the deletion is canceled and an appropriate error is returned.
//
// It returns a success message upon successful deletion, or an error if deletion fails
// or the user cancels the operation.
func handleDelete(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("expected task ID for deletion")
	}

	id := strings.ToLower(strings.TrimSpace(args[1]))

	fmt.Printf("Are you sure you want to delete task '%s'? [y/N]: ", id)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	input = strings.ToLower(strings.TrimSpace(input))
	if input != "y" && input != "yes" {
		return "", fmt.Errorf("task deletion cancelled")
	}

	if err := svc.DeleteTask(id); err != nil {
		return "", fmt.Errorf("failed to delete task: %w", err)
	}

	return fmt.Sprintf("Task %s deleted successfully", id), nil
}

func handleMarkStatus(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: task mark-<status> <task-id>")
		return
	}

	statusInput := args[0]
	taskID := args[1]

	status, err := task.ParseAndValidateStatus(statusInput)
	if err != nil {
		fmt.Println("invalid status:", err)
		return
	}

	if err := svc.UpdateTaskStatus(taskID, status); err != nil {
		fmt.Println("failed to update task status:", err)
		return
	}

	fmt.Println("Task status updated successfully")
}

func handleUpdateDescription(args []string) {
	if len(args) < 3 {
		fmt.Println("usage: task update <task-id> <new description>")
		return
	}

	taskID := args[1]
	newDescription := strings.Join(args[2:], " ")

	if err := svc.UpdateTaskDescription(newDescription, taskID); err != nil {
		fmt.Println("failed to update task description:", err)
		return
	}

	fmt.Println("Task description updated successfully!")
}
