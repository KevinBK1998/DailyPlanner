package store

import (
	"errors"
	"sort"
	"sync"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/models"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskStore struct {
	mu     sync.Mutex
	tasks  map[int]models.Task
	nextID int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *TaskStore) List() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]models.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		out = append(out, t)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].ID < out[j].ID
	})

	return out
}

func (s *TaskStore) Add(title string) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := models.Task{
		ID:     s.nextID,
		Title:  title,
		Status: "pending",
	}
	s.tasks[s.nextID] = task
	s.nextID++

	return task
}

func (s *TaskStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(s.tasks, id)
	return nil
}

func (s *TaskStore) Complete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return ErrTaskNotFound
	}

	task.Status = "completed"
	s.tasks[id] = task
	return nil
}
