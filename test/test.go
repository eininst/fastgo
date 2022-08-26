package main

import "fmt"

type UserService interface {
	Addx(name string) error
}

type UserServiceImpl struct {
}

func (us *UserServiceImpl) Addx(name string) error {
	fmt.Println(name)
	return nil
}

func main() {
	var us UserService

	us = &UserServiceImpl{}

	_ = us.Addx("hello")
}
