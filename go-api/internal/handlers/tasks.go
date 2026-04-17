package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/store"
)

type createTaskRequest struct {
	Title string `json:"title"`
}

func HandleTasks(taskStore *store.TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

		if len(parts) == 2 && parts[0] == "tasks" && r.Method == http.MethodDelete {
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				http.Error(w, "invalid task id", http.StatusBadRequest)
				return
			}

			if err := taskStore.Delete(id); err != nil {
				http.Error(w, "task not found", http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusNoContent)
			return
		}

		if len(parts) == 3 && parts[0] == "tasks" && parts[2] == "complete" && r.Method == http.MethodPut {
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				http.Error(w, "invalid task id", http.StatusBadRequest)
				return
			}

			if err := taskStore.Complete(id); err != nil {
				http.Error(w, "task not found", http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusNoContent)
			return
		}

		if !(len(parts) == 1 && parts[0] == "tasks") {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(taskStore.List())

		case http.MethodPost:
			var req createTaskRequest

			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "invalid json body", http.StatusBadRequest)
				return
			}

			title := strings.TrimSpace(req.Title)
			if title == "" {
				http.Error(w, "title is required", http.StatusBadRequest)
				return
			}

			task := taskStore.Add(title)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(task)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
