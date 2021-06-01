package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(pingFunc)
	log.Fatal(http.ListenAndServe(":80", handler))
}

func pingFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
