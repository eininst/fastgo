package user

import (
	"github.com/eininst/flog"
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
	flog.Info("add123")
}

func (us *userService) Update() {

}
