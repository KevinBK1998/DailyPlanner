package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/KevinBK1998/dailyplanner/go-api/internal/handlers"
	"github.com/KevinBK1998/dailyplanner/go-api/internal/store"
)

func main() {
	taskStore, err := store.NewTaskStore()
	if err != nil {
		fmt.Printf("error creating task store: %v\n", err)
		return
	}
	defer taskStore.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})
	mux.HandleFunc("/tasks", handlers.HandleTasks(taskStore))
	mux.HandleFunc("/tasks/", handlers.HandleTasks(taskStore))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server starting on :8080")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("server shutdown error: %v", err)
		}
	case err := <-errCh:
		if err != nil && err != http.ErrServerClosed {
			log.Printf("server failed to start: %v", err)
		}
	}
}
