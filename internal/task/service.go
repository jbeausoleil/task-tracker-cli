package task

import (
	"fmt"
	"task-tracker-cli/internal/idgen"
	"time"
)

// Service provides methods for managing tasks, including creating, listing, and deleting tasks.
type Service struct {
	store *Store
}

// NewService initializes and returns a new Service instance using the provided Store.
func NewService(store *Store) *Service {
	return &Service{
		store: store,
	}
}

// CreateTask creates a new task with the given description,
// sets its default status and timestamps, saves it to the store, and returns the created task.
func (s *Service) CreateTask(desc string) (Task, error) {
	uuid := idgen.GenerateID()
	task := Task{
		Id:          uuid,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.store.AppendTask(task); err != nil {
		return Task{}, fmt.Errorf("failed to append task: %w", err)
	}

	return task, nil
}

// ListTasks retrieves all tasks, optionally filtering by status.
// If filter is an empty string, it returns all tasks.
func (s *Service) ListTasks(filter string) []Task {
	var result []Task
	for _, task := range s.store.tasks {
		if filter == "" || string(task.Status) == filter {
			result = append(result, task)
		}
	}
	return result
}

// DeleteTask deletes the task with the specified ID from the store and saves the updated list.
// It returns an error if the task is not found or saving to file fails.
func (s *Service) DeleteTask(id string) error {
	if err := s.store.DeleteTaskById(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	if err := s.store.SaveToFile(); err != nil {
		return fmt.Errorf("failed to save tasks to file: %w", err)
	}

	return nil
}
