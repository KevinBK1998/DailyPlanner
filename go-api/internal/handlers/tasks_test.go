package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/models"
	"github.com/KevinBK1998/dailyplanner/go-api/internal/store"
)

func TestHandleTasks_CreateThenList(t *testing.T) {
	s := store.NewTaskStore()
	h := HandleTasks(s)

	createReq := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Learn tests"}`))
	createReq.Header.Set("Content-Type", "application/json")
	createRes := httptest.NewRecorder()
	h(createRes, createReq)

	if createRes.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, createRes.Code)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	listRes := httptest.NewRecorder()
	h(listRes, listReq)

	if listRes.Code != http.StatusOK {
		t.Fatalf("expected %d, got %d", http.StatusOK, listRes.Code)
	}

	var tasks []models.Task
	if err := json.NewDecoder(listRes.Body).Decode(&tasks); err != nil {
		t.Fatalf("failed to decode list response: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "Learn tests" {
		t.Fatalf("expected title %q, got %q", "Learn tests", tasks[0].Title)
	}
}

func TestHandleTasks_Create_InvalidJSON(t *testing.T) {
    s := store.NewTaskStore()
    h := HandleTasks(s)

    req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader("{bad json"))
    req.Header.Set("Content-Type", "application/json")
    res := httptest.NewRecorder()
    h(res, req)

    if res.Code != http.StatusBadRequest {
        t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
    }
}
