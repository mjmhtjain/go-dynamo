package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/repo"
)

var userDetailRepo repo.UserDetailRepo

func UserHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Path[len("/user/"):]

	log.Printf("userId: %v", userId)

	if userDetailRepo == nil {
		log.Println("NewUserDetailRepo created .. ")
		userDetailRepo = repo.NewUserDetailRepo()
	}
	resp, err := userDetailRepo.FindById(userId)

	if err != nil {
		log.Fatal("something went wrong")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
}
