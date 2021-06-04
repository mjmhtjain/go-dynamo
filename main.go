package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mjmhtjain/go-dynamo/src/router"
)

const (
	DEFAULT_ADDR string = ":80"
)

func main() {
	godotenv.Load()
	r := router.Router()
	log.Fatal(http.ListenAndServe(DEFAULT_ADDR, r))
}
