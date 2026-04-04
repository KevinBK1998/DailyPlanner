package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/models"
)

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []models.Task{
		{ID: 1, Title: "Learn Go", Status: "pending"},
		{ID: 2, Title: "Build REST API", Status: "pending"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
