package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/store"
)

func HandleTasks(taskStore *store.TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(taskStore.List())
	}
}
