package postgres

import (
	"fmt"

	"golang.source-fellows.com/samples/applicationx"
)

var _ applicationx.UserService = &UserService{}

type UserService struct{}

func (us *UserService) CreateUser(user *applicationx.User) error {
	fmt.Println("Create User in Postgres Service")
	return nil
}

func (us *UserService) DeleteUser(id int) error {
	return nil
}

func (us *UserService) ReadUser(id int) (*applicationx.User, error) {
	return nil, nil
}
