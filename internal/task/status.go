package task

import (
	"fmt"
	"strings"
)

// validStatuses defines the allowed status strings for tasks.
var validStatuses = map[string]Status{
	"todo":        StatusTodo,
	"in_progress": StatusInProgress,
	"completed":   StatusCompleted,
}

// IsValidStatus checks if a given string is a valid Status value.
func isValidStatus(s string) bool {
	_, exists := validStatuses[s]
	return exists
}

// parseStatus converts a string to a Status if valid, or returns an error.
func parseStatus(s string) (Status, error) {
	status, exists := validStatuses[s]
	if !exists {
		return "", fmt.Errorf("invalid status: %s", s)
	}
	return status, nil
}

func parseMarkCommand(input string) (string, error) {
	input = strings.ToLower(strings.TrimSpace(input))

	if !strings.HasPrefix(input, "mark-") {
		return "", fmt.Errorf("invalid command: expected prefix 'mark-'")
	}

	// Strip the "mark-" part
	statusPart := strings.TrimPrefix(input, "mark-")

	// Replace internal dashes with underscores
	statusPart = strings.ReplaceAll(statusPart, "-", "_")

	return statusPart, nil
}

func ParseAndValidateStatus(input string) (Status, error) {
	statusInput, err := parseMarkCommand(input)
	if err != nil {
		return "", fmt.Errorf("invalid mark command: %w", err)
	}

	if !isValidStatus(statusInput) {
		return "", fmt.Errorf("invalid status: %s", statusInput)
	}

	status, err := parseStatus(statusInput)
	if err != nil {
		return "", fmt.Errorf("failed to parse status: %w", err)
	}

	return status, nil
}
