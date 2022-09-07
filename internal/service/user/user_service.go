package user

import (
	"fastgo/internal/types"
	"github.com/eininst/rlock"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserService interface {
	Add() error
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

func (us *userService) Add() error {
	panic("weweww")
	return types.NewServiceError("my name is error")
}

func (us *userService) Update() {

}
