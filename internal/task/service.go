package task

import (
	"task-tracker-cli/internal/idgen"
	"time"
)

// Service provides methods to manage tasks by interacting with the underlying data store.
type Service struct {
	store *Store
}

// NewService initializes and returns a new Service instance with the provided Store.
func NewService(store *Store) *Service {
	return &Service{
		store: store,
	}
}

// CreateTask adds a new task with the provided description to the store and returns the created Task object.
func (s *Service) CreateTask(desc string) (Task, error) {
	uuid := idgen.GenerateID()
	task := Task{
		Id:          uuid,
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := s.store.AppendTask(task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
