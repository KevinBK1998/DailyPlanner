package store

import "testing"

func TestAddAndList(t *testing.T) {
	s, err := NewTaskStore()
	if err != nil {
		t.Fatalf("failed to create TaskStore: %v", err)
	}
	defer s.tasks.Close()
	if _, err := s.tasks.Exec("DELETE FROM tasks"); err != nil {
		t.Fatalf("failed to clear tasks table: %v", err)
	}
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

func TestDeleteAndCompleteErrors(t *testing.T) {
	s, err := NewTaskStore()
	if err != nil {
		t.Fatalf("failed to create TaskStore: %v", err)
	}
	defer s.tasks.Close()
	if _, err := s.tasks.Exec("DELETE FROM tasks"); err != nil {
		t.Fatalf("failed to clear tasks table: %v", err)
	}
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
