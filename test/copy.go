package main

import (
	"github.com/eininst/flog"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/jinzhu/copier"
)

type User struct {
	Name   string
	Role   string
	Age    int32
	Salary int
}

type UserDTO struct {
	Name string
	Role string
	Age  int64
}

func main() {
	var (
		user = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000}
	)
	var udto UserDTO
	copier.Copy(&udto, &user)
	flog.Info(udto)

	flog.Info(utils.UUIDv4())
	flog.Info(utils.UUIDv4())
}
