package main

type UserService interface {
	Save()
}

var _ UserService = &MyUserService{}

type MyUserService struct {
}

func (us *MyUserService) Save() {

}

func main() {

}
