package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mjmhtjain/go-dynamo/src/customError"
	"github.com/mjmhtjain/go-dynamo/src/model"
	"github.com/mjmhtjain/go-dynamo/src/repo"
)

var userDetailRepo repo.UserDetailRepo

func UserHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		findUser(w, r)
	case http.MethodPost:
		saveUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}

}

func findUser(w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/user/")

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

func saveUser(w http.ResponseWriter, r *http.Request) {
	user := new(model.UserDetail)
	json.NewDecoder(r.Body).Decode(&user)

	log.Printf("user: %v", user)

	if userDetailRepo == nil {
		log.Println("NewUserDetailRepo created .. ")
		userDetailRepo = repo.NewUserDetailRepo()
	}
	err := userDetailRepo.Save(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
