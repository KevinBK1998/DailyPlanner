package store

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/models"
	_ "modernc.org/sqlite"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskStore struct {
	mu    sync.Mutex
	tasks *sql.DB
}

func NewTaskStore() (*TaskStore, error) {
	return NewTaskStoreWithPath("tasks.db")
}

func NewTaskStoreWithPath(dbPath string) (*TaskStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT, status TEXT)`)
	if err != nil {
		_ = db.Close()
		return nil, err
	}
	return &TaskStore{
		tasks: db,
	}, nil
}

func (s *TaskStore) Close() error {
	return s.tasks.Close()
}

func (s *TaskStore) List() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	rows, err := s.tasks.Query(`SELECT * FROM tasks ORDER BY id`)
	if err != nil {
		fmt.Printf("error querying tasks: %v\n", err)
		return nil
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Status); err != nil {
			fmt.Printf("error scanning task: %v\n", err)
			return nil
		}
		tasks = append(tasks, t)
	}
	return tasks
}

func (s *TaskStore) Add(title string) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := s.tasks.Exec(`INSERT INTO tasks (title, status) VALUES (?, ?)`, title, "pending")
	if err != nil {
		fmt.Printf("error inserting task: %v\n", err)
		return models.Task{}
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("error getting last insert id: %v\n", err)
		return models.Task{}
	}
	return models.Task{
		ID:     int(id),
		Title:  title,
		Status: "pending",
	}
}

func (s *TaskStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := s.tasks.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		fmt.Printf("error deleting task: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("error getting rows affected: %v\n", err)
		return err
	}
	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (s *TaskStore) Complete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := s.tasks.Exec(`UPDATE tasks SET status = ? WHERE id = ?`, "completed", id)
	if err != nil {
		fmt.Printf("error updating task: %v\n", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("error getting rows affected: %v\n", err)
		return err
	}
	if rowsAffected == 0 {
		return ErrTaskNotFound
	}

	return nil
}
