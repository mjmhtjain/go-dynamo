package main

import (
	"log"
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/router"
)

const (
	DEFAULT_ADDR string = ":80"
)

func main() {
	r := router.Router()
	log.Fatal(http.ListenAndServe(DEFAULT_ADDR, r))
}
