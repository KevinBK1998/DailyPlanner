package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{ID: 1, Title: "Learn Go", Status: "pending"},
		{ID: 2, Title: "Build REST API", Status: "pending"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})
	http.HandleFunc("/tasks", handleTasks)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
