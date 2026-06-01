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
	first, err := s.Add("First Task")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}
	second, err := s.Add("Second Task")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}

	if first.ID != 1 || second.ID != 2 {
		t.Fatalf("expectd IDs 1 and 2, got %d and %d", first.ID, second.ID)
	}

	tasks, err := s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Title != "First Task" || tasks[1].Title != "Second Task" {
		t.Fatalf("unexpected task order/titles: %+v", tasks)
	}
}

func TestAddAndUpdate(t *testing.T) {
	s := newTestTaskStore(t)
	task, err := s.Add("First Task")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}

	if task.ID != 1 {
		t.Fatalf("expectd ID 1, got %d", task.ID)
	}

	tasks, err := s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "First Task" {
		t.Fatalf("unexpected task order/titles: %+v", tasks)
	}

	if err := s.Update(task.ID, "Updated Task"); err != nil {
		t.Fatalf("expected update to succeed, got %v", err)
	}
	tasks, err = s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Title != "Updated Task" {
		t.Fatalf("expected title %q, got %q", "Updated Task", tasks[0].Title)
	}

	if err := s.Update(999, "Nonexistent Task"); err != ErrTaskNotFound {
		t.Fatalf("expected ErrTaskNotFound, got %v", err)
	}
}

func TestDeleteAndList(t *testing.T) {
	s := newTestTaskStore(t)
	first, err := s.Add("First Task")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}
	second, err := s.Add("Second Task")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}
	third, err := s.Add("Third Task")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}

	if first.ID != 1 || second.ID != 2 || third.ID != 3 {
		t.Fatalf("expectd IDs 1, 2, and 3, got %d, %d, and %d", first.ID, second.ID, third.ID)
	}

	if err := s.Delete(second.ID); err != nil {
		t.Fatalf("expected delete to succeed, got %v", err)
	}

	tasks, err := s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0].Title != "First Task" || tasks[1].Title != "Third Task" {
		t.Fatalf("unexpected task order/titles: %+v", tasks)
	}
}

func TestDeleteAndCompleteErrors(t *testing.T) {
	s := newTestTaskStore(t)
	task, err := s.Add("Only Task")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}

	if err := s.Complete(task.ID); err != nil {
		t.Fatalf("expected complete to succeed, got %v", err)
	}

	tasks, err := s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	if tasks[0].Status != "completed" {
		t.Fatalf("expected task to be completed, got %s", tasks[0].Status)
	}

	if err := s.Delete(task.ID); err != nil {
		t.Fatalf("expected delete to succeed, got %v", err)
	}

	tasks, err = s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
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
				task, err := s.Add(fmt.Sprintf("w%d-task-%d", worker, i))
				if err != nil {
					errCh <- fmt.Errorf("failed to add task for worker=%d: i=%d: %v", worker, i, err)
					continue
				}
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

	tasks, err := s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	want := workers * perWorker
	if len(tasks) != want {
		t.Fatalf("expected %d tasks, got %d", want, len(tasks))
	}
}

func TestAddAndList_Concurrent(t *testing.T) {
	s := newTestTaskStore(t)

	const writerWorkers = 10
	const readerWorkers = 10
	const perWriter = 20

	var wg sync.WaitGroup
	errCh := make(chan error, writerWorkers*perWriter)

	for w := 0; w < writerWorkers; w++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for i := 0; i < perWriter; i++ {
				task, err := s.Add(fmt.Sprintf("writer-%d-task-%d", worker, i))
				if err != nil {
					errCh <- fmt.Errorf("failed to add task for writer=%d: i=%d: %v", worker, i, err)
					continue
				}
				if task.ID == 0 {
					errCh <- fmt.Errorf("failed to add task for writer=%d: i=%d", worker, i)
				}
			}
		}(w)
	}

	for r := 0; r < readerWorkers; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < perWriter; i++ {
				if _, err := s.List(); err != nil {
					errCh <- fmt.Errorf("failed to list tasks: %v", err)
				}
			}
		}()
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.Fatal(err)
	}

	tasks, err := s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	want := writerWorkers * perWriter
	if len(tasks) != want {
		t.Fatalf("expected %d tasks, got %d", want, len(tasks))
	}
}

func TestAddAndList_ClosedStoreErrors(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "tasks.db")
	s, err := NewTaskStoreWithPath(dbPath)
	if err != nil {
		t.Fatalf("failed to create TaskStore: %v", err)
	}

	if err := s.Close(); err != nil {
		t.Fatalf("failed to close TaskStore: %v", err)
	}

	if _, err := s.Add("Should Fail"); err == nil {
		t.Fatal("expected add to fail on closed store")
	}

	if _, err := s.List(); err == nil {
		t.Fatal("expected list to fail on closed store")
	}
}
