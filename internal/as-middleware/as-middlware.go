package internal

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Health struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Uuid      string `json:"uuid"`
}

type Uuid struct{}

func GuidMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		uuid := uuid.New().String()
		r = r.WithContext(context.WithValue(r.Context(), Uuid{}, uuid))
		next.ServeHTTP(rw, r)
	})
}
