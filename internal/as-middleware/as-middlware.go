package internal

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	tools "github.com/terrytay/go-context/tools/json"
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

func HealthCheck(rw http.ResponseWriter, r *http.Request) {
	uuid := r.Context().Value(Uuid{}) // uuid from middlware

	response := Health{
		Message:   "OK",
		Timestamp: time.Now().UnixMicro(),
		Uuid:      uuid.(string),
	}

	rw.WriteHeader(http.StatusOK)

	tools.ToJSON(response, rw)
}
