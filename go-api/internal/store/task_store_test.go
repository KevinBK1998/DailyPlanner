package store

import (
	"fmt"
	"path/filepath"
	"sync"
	"testing"
)

func newTestTaskStore(t *testing.T) *TaskStore {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "tasks.db")
	s, err := NewTaskStoreWithPath(dbPath)
	if err != nil {
		t.Fatalf("failed to create TaskStore: %v", err)
	}

	t.Cleanup(func() {
		if err := s.Close(); err != nil {
			t.Fatalf("failed to close TaskStore: %v", err)
		}
	})

	return s
}

func TestAddAndList(t *testing.T) {
	s := newTestTaskStore(t)
	first := s.Add("First Task")
	second := s.Add("Second Task")

	if first.ID != 1 || second.ID != 2 {
		t.Fatalf("expectd IDs 1 and 2, got %d and %d", first.ID, second.ID)
	}

	tasks := s.List()
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Title != "First Task" || tasks[1].Title != "Second Task" {
		t.Fatalf("unexpected task order/titles: %+v", tasks)
	}
}

func TestDeleteAndList(t *testing.T) {
	s := newTestTaskStore(t)
	first := s.Add("First Task")
	second := s.Add("Second Task")
	third := s.Add("Third Task")

	if first.ID != 1 || second.ID != 2 || third.ID != 3 {
		t.Fatalf("expectd IDs 1, 2, and 3, got %d, %d, and %d", first.ID, second.ID, third.ID)
	}

	if err := s.Delete(second.ID); err != nil {
		t.Fatalf("expected delete to succeed, got %v", err)
	}

	tasks := s.List()
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Title != "First Task" || tasks[1].Title != "Third Task" {
		t.Fatalf("unexpected task order/titles: %+v", tasks)
	}
}

func TestDeleteAndCompleteErrors(t *testing.T) {
	s := newTestTaskStore(t)
	task := s.Add("Only Task")

	if err := s.Complete(task.ID); err != nil {
		t.Fatalf("expected complete to succeed, got %v", err)
	}

	tasks := s.List()
	if tasks[0].Status != "completed" {
		t.Fatalf("expected task to be completed, got %s", tasks[0].Status)
	}

	if err := s.Delete(task.ID); err != nil {
		t.Fatalf("expected delete to succeed, got %v", err)
	}

	tasks = s.List()
	if len(tasks) != 0 {
		t.Fatalf("expected 0 tasks after delete, got %d", len(tasks))
	}

	if err := s.Delete(task.ID); err != ErrTaskNotFound {
		t.Fatalf("expected ErrTaskNotFound on second delete, got %v", err)
	}

	if err := s.Complete(task.ID); err != ErrTaskNotFound {
		t.Fatalf("expected ErrTaskNotFound on complete after delete, got %v", err)
	}
}

func TestAdd_Concurrent(t *testing.T) {
	s := newTestTaskStore(t)
	const workers = 20
	const perWorker = 10
	var wg sync.WaitGroup
	errCh := make(chan error, workers*perWorker)

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for i := 0; i < perWorker; i++ {
				task := s.Add(fmt.Sprintf("w%d-task-%d", worker, i))
				if task.ID == 0 {
					errCh <- fmt.Errorf("failed to add task for worker=%d: i=%d", worker, i)
				}
			}
		}(w)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Fatal(err)
	}

	tasks := s.List()
	want := workers * perWorker
	if len(tasks) != want {
		t.Fatalf("expected %d tasks, got %d", want, len(tasks))
	}
}
