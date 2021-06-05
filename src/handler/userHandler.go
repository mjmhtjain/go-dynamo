package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/customError"
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
		var recNotFound *customError.RecordNotFoundError
		log.Printf("UserDetail returned an err: %v", err)

		if errors.As(err, &recNotFound) {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
}
