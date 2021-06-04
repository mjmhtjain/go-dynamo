package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mjmhtjain/go-dynamo/src/repo"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Path[len("/user/"):]

	log.Printf("userId: %v", userId)

	userDetailRepo := repo.NewUserDetailRepo()

	resp, err := userDetailRepo.FindById(userId)

	if err != nil {
		log.Fatal("something went wrong")
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, resp)
}
