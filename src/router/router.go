package router

import (
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/handler"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/health", http.HandlerFunc(handler.HealthHandler))
	router.Handle("/user/", http.HandlerFunc(handler.UserHandler))

	return router
}
