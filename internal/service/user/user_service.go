package user

import (
	"fastgo/internal/common/serr"
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
	return serr.NewServiceError("my name is error")
}

func (us *userService) Update() {

}
