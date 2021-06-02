package repo

import "github.com/mjmhtjain/go-dynamo/src/model"

type UserDetailRepo interface {
	FindUserById(id string) *model.UserDetail
}
