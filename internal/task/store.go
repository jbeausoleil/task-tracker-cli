package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-tracker-cli/internal/idgen"
	"time"
)

// Status represents the state of a task or process, typically used to track progress or completion stages.
type Status string

// StatusTodo indicates a task is yet to be started.
// StatusInProgress indicates a task is currently being worked on.
// StatusCompleted indicates a task has been finished.
const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusCompleted  Status = "completed"
)

// taskWrapper is a container type that wraps a list of Task objects for serialization or data manipulation purposes.
type taskWrapper struct {
	Tasks []Task `json:"tasks"`
}

// Store represents a data structure for managing Task items and their persistence to a file.
type Store struct {
	tasks    []Task
	filePath string
}

// fileExists checks if the specified file path exists and returns true if it exists, or false otherwise.
func fileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

// createEmptyTaskFile creates a new JSON file at the specified path and populates it with an empty list of tasks.
// It panics if the file cannot be created or written to.
func createEmptyTaskFile(path string) {
	emptyData, _ := json.MarshalIndent([]Task{}, "", "  ")
	if err := os.WriteFile(path, emptyData, 0644); err != nil {
		panic(fmt.Errorf("failed to create tasks.json file: %w", err))
	}
}

// loadTasksFromFile reads tasks from a JSON file specified by the given path and returns a slice of Task or an error.
func loadTasksFromFile(path string) ([]Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var wrapper taskWrapper
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Tasks, nil
}

// NewStore initializes and returns a new Store instance with tasks loaded from a JSON file or an empty file if not present.
func NewStore() *Store {
	path := "db/task.json"
	if ok := fileExists(path); !ok {
		createEmptyTaskFile(path)
	}

	tasks, err := loadTasksFromFile(path)
	if err != nil {
		fmt.Println("failed to load tasks from file: ", err)
		os.Exit(1)
	}

	return &Store{
		tasks:    tasks,
		filePath: path,
	}
}

// AddTask adds a new task with the given description to the store, saves it to a file, and returns the created Task.
func (s *Store) AddTask(desc string) Task {
	uuid := idgen.GenerateID()
	task := Task{
		Id:          uuid,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	s.tasks = append(s.tasks, task)
	err := s.SaveToFile()
	if err != nil {
		panic(err)
	}
	return task
}

// SaveToFile saves the current tasks in the Store to a file in JSON format with indentation. Returns an error if failed.
func (s *Store) SaveToFile() error {
	wrapped := taskWrapper{Tasks: s.tasks}
	data, err := json.MarshalIndent(wrapped, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}
