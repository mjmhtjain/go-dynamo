package main

import (
	"log"
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/router"
)

const (
	DEFAULT_PORT string = ":80"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(DEFAULT_PORT, r))
}
