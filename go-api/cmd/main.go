package main

import (
	"fmt"
	"net/http"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/handlers"
	"github.com/KevinBK1998/dailyplanner/go-api/internal/store"
)

func main() {
	taskStore := store.NewTaskStore()
	taskStore.Add("Learn Go")
	taskStore.Add("Build REST API")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})
	http.HandleFunc("/tasks", handlers.HandleTasks(taskStore))

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
