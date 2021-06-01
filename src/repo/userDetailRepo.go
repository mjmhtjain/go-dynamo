package repo

type UserDetailRepo interface {
	FindUserById(id string)
}
