package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Status represents the state of a task, used to track its progress or completion stage.
type Status string

const (
	// StatusTodo indicates the task is yet to be started.
	StatusTodo Status = "todo"

	// StatusInProgress indicates the task is currently being worked on.
	StatusInProgress Status = "in_progress"

	// StatusCompleted indicates the task has been finished.
	StatusCompleted Status = "completed"
)

// taskWrapper is a container for serializing and deserializing a list of tasks.
type taskWrapper struct {
	Tasks []Task `json:"tasks"`
}

// Store manages a collection of tasks and their persistence to a file.
type Store struct {
	tasks    []Task
	filePath string
}

// fileExists returns true if a file exists at the given path.
func fileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// createEmptyTaskFile creates a new JSON file at the specified path with an empty list of tasks.
// It panics if the file cannot be created or written.
func createEmptyTaskFile(path string) {
	var wrapper taskWrapper
	data, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		panic(fmt.Errorf("failed to marshal empty task wrapper: %w", err))
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		panic(fmt.Errorf("failed to create tasks file: %w", err))
	}
}

// NewStore initializes and returns a new Store instance.
// It loads existing tasks from a JSON file, creating an empty file if it does not exist.
func NewStore() *Store {
	const path = "db/task.json"

	if !fileExists(path) {
		createEmptyTaskFile(path)
	}

	tasks, err := LoadTasksFromFile(path)
	if err != nil {
		fmt.Println("failed to load tasks from file:", err)
		os.Exit(1)
	}

	return &Store{
		tasks:    tasks,
		filePath: path,
	}
}

// AppendTask adds a new task to the store and immediately saves the updated list to the file.
func (s *Store) AppendTask(task Task) error {
	s.tasks = append(s.tasks, task)
	return s.SaveToFile()
}

// SaveToFile writes the current tasks to the store's file in JSON format.
// It overwrites the existing file contents.
func (s *Store) SaveToFile() error {
	wrapped := taskWrapper{Tasks: s.tasks}
	data, err := json.MarshalIndent(wrapped, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}
	return os.WriteFile(s.filePath, data, 0644)
}

// LoadTasksFromFile reads tasks from the given JSON file path and returns the list of tasks.
// It returns an error if the file cannot be read or parsed.
func LoadTasksFromFile(path string) ([]Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var wrapper taskWrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	return wrapper.Tasks, nil
}

// DeleteTaskById removes a task from the store by its ID.
// It updates the in-memory task list but does not automatically save to file.
// Returns an error if the task ID does not exist.
func (s *Store) DeleteTaskById(id string) error {
	found := false
	tasks := s.tasks

	i := 0
	for _, task := range tasks {
		if task.Id == id {
			found = true
			continue
		}
		tasks[i] = task
		i++
	}

	if !found {
		return fmt.Errorf("task %s not found", id)
	}

	s.tasks = tasks[:i]
	return nil
}
