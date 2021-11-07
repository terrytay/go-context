package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	internal "github.com/terrytay/go-context/internal/as-middleware"
	tools "github.com/terrytay/go-context/tools/json"
)

func main() {
	l := log.New(os.Stdout, "web", log.LstdFlags)
	r := chi.NewRouter()

	r.Use(internal.GuidMiddleware)
	r.HandleFunc("/health", healthCheck)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		IdleTimeout:  120 * time.Second, // max time for connections using TCP keep-alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("received termination. graceful shutdown...", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

func healthCheck(rw http.ResponseWriter, r *http.Request) {
	uuid := r.Context().Value(internal.Uuid{})

	response := internal.Health{
		Message:   "OK",
		Timestamp: time.Now().UnixMicro(),
		Uuid:      uuid.(string),
	}

	rw.WriteHeader(http.StatusOK)

	tools.ToJSON(response, rw)
}
