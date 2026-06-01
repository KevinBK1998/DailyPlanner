package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/models"
	"github.com/KevinBK1998/dailyplanner/go-api/internal/store"
)

func newTestTaskStore(t *testing.T) *store.TaskStore {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "tasks.db")
	s, err := store.NewTaskStoreWithPath(dbPath)
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

func newManualCloseTestTaskStore(t *testing.T) *store.TaskStore {
	t.Helper()

	dbPath := filepath.Join(t.TempDir(), "tasks.db")
	s, err := store.NewTaskStoreWithPath(dbPath)
	if err != nil {
		t.Fatalf("failed to create TaskStore: %v", err)
	}

	return s
}

func TestHandleTasks_CreateThenList(t *testing.T) {
	s := newTestTaskStore(t)
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
	s := newTestTaskStore(t)
	h := HandleTasks(s)

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader("{bad json"))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
	}
}

func TestHandleTasks_Delete_NonexistentId(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodDelete, "/tasks/999", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("expected %d, got %d", http.StatusNotFound, res.Code)
	}
}

func TestHandleTasks_Delete_InvalidId(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodDelete, "/tasks/abc", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
	}
}
func TestHandleTasks_Complete_NonexistentId(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodPut, "/tasks/999/complete", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("expected %d, got %d", http.StatusNotFound, res.Code)
	}
}

func TestHandleTasks_Complete_InvalidId(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodPut, "/tasks/abc/complete", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
	}
}

func TestHandleTasks_Patch_InvalidId(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodPatch, "/tasks/abc", strings.NewReader(`{"title":"Learn test"}`))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
	}
}

func TestHandleTasks_Patch_NonexistentId(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodPatch, "/tasks/999", strings.NewReader(`{"title":"Learn test"}`))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("expected %d, got %d", http.StatusNotFound, res.Code)
	}
}

func TestHandleTasks_Patch_InvalidJSON(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)

	req := httptest.NewRequest(http.MethodPatch, "/tasks/999", strings.NewReader("{bad json"))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, got %d", http.StatusBadRequest, res.Code)
	}
}

func TestHandleTasks_Create_Patch(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)

	createReq := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Learn test"}`))
	createReq.Header.Set("Content-Type", "application/json")
	createRes := httptest.NewRecorder()
	h(createRes, createReq)

	if createRes.Code != http.StatusCreated {
		t.Fatalf("expected %d, got %d", http.StatusCreated, createRes.Code)
	}

	patchReq := httptest.NewRequest(http.MethodPatch, "/tasks/1", strings.NewReader(`{"title":"Learn tests"}`))
	patchReq.Header.Set("Content-Type", "application/json")
	patchRes := httptest.NewRecorder()
	h(patchRes, patchReq)

	if patchRes.Code != http.StatusNoContent {
		t.Fatalf("expected %d, got %d", http.StatusNoContent, patchRes.Code)
	}

	tasks, err := s.List()
	if err != nil {
		t.Fatalf("expected list to succeed, got %v", err)
	}
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Title != "Learn tests" {
		t.Fatalf("expected title %q, got %q", "Learn tests", tasks[0].Title)
	}
}

func TestHandleTasks_List_StoreErrorReturnsInternalServerError(t *testing.T) {
	s := newManualCloseTestTaskStore(t)
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close TaskStore: %v", err)
	}
	h := HandleTasks(s)

	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, res.Code)
	}
}

func TestHandleTasks_Create_StoreErrorReturnsInternalServerError(t *testing.T) {
	s := newManualCloseTestTaskStore(t)
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close TaskStore: %v", err)
	}
	h := HandleTasks(s)

	req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Learn tests"}`))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, res.Code)
	}
}

func TestHandleTasks_Patch_StoreErrorReturnsInternalServerError(t *testing.T) {
	s := newManualCloseTestTaskStore(t)
	task, err := s.Add("Learn tests")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close TaskStore: %v", err)
	}
	h := HandleTasks(s)

	req := httptest.NewRequest(http.MethodPatch, "/tasks/"+strconv.Itoa(task.ID), strings.NewReader(`{"title":"Updated"}`))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, res.Code)
	}
}

func TestHandleTasks_Delete_StoreErrorReturnsInternalServerError(t *testing.T) {
	s := newManualCloseTestTaskStore(t)
	task, err := s.Add("Learn tests")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close TaskStore: %v", err)
	}
	h := HandleTasks(s)

	req := httptest.NewRequest(http.MethodDelete, "/tasks/"+strconv.Itoa(task.ID), nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, res.Code)
	}
}

func TestHandleTasks_Complete_StoreErrorReturnsInternalServerError(t *testing.T) {
	s := newManualCloseTestTaskStore(t)
	task, err := s.Add("Learn tests")
	if err != nil {
		t.Fatalf("expected add to succeed, got %v", err)
	}
	if err := s.Close(); err != nil {
		t.Fatalf("failed to close TaskStore: %v", err)
	}
	h := HandleTasks(s)

	req := httptest.NewRequest(http.MethodPut, "/tasks/"+strconv.Itoa(task.ID)+"/complete", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, res.Code)
	}
}

func TestHandleTasks_UnsupportedMethod(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodPatch, "/tasks", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected %d, got %d", http.StatusMethodNotAllowed, res.Code)
	}
}
func TestHandleTasks_UnsupportedPath(t *testing.T) {
	s := newTestTaskStore(t)
	h := HandleTasks(s)
	req := httptest.NewRequest(http.MethodGet, "/tasks/1/path", nil)
	res := httptest.NewRecorder()
	h(res, req)

	if res.Code != http.StatusNotFound {
		t.Fatalf("expected %d, got %d", http.StatusNotFound, res.Code)
	}
}
