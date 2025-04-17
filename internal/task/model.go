package task

import "time"

// Task represents a task with unique ID, description, status, and timestamps for creation and updates.
type Task struct {
	Id          string    `json:"id" :"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
