package user

import (
	"fmt"
	"github.com/eininst/rlock"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserService interface {
	Add()
	Update()
}
type userService struct {
	DB          *gorm.DB      `inject:""`
	RedisClient *redis.Client `inject:""`
	Rlock       *rlock.Rlock  `inject:""`
}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) Add() {
	fmt.Println(us.RedisClient)
	fmt.Println(us.DB)
	fmt.Println(us.Rlock)
	fmt.Println("add")
}
func (us *userService) Update() {

}
