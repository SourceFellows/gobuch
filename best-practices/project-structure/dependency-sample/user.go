//go:generate mockgen -source=user.go -package mocks -destination mocks/user_service.go

package applicationx

type User struct {
	ID   int
	name string
}

type UserService interface {
	CreateUser(u *User) error
	ReadUser(id int) (*User, error)
	DeleteUser(id int) error
}
