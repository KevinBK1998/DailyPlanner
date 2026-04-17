package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/store"
)

type createTaskRequest struct {
	Title string `json:"title"`
}

func HandleTasks(taskStore *store.TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
