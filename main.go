package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/router"
)

func main() {
	// handler := http.HandlerFunc(pingFunc)
	r := router.Router()
	log.Fatal(http.ListenAndServe(":80", r))
}

func pingFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
