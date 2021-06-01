package router

import (
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/health"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/health", http.HandlerFunc(health.HealthHandler))

	return router
}
