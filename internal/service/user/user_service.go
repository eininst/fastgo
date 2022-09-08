package user

import (
	"fastgo/internal/code"
	"fastgo/internal/common/serr"
	"fastgo/internal/model"
	"fmt"
	"github.com/eininst/rlock"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserService interface {
	Get(id int64) (*UserDTO, error)
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

func (us *userService) Get(id int64) (*UserDTO, error) {
	var u model.User
	us.DB.First(&u, id)
	if u.Id == 0 {
		msg := fmt.Sprintf("user is not found by %v", id)
		return nil, serr.NewServiceError(msg, code.ERROR_DATA_NOT_FOUND)
	}

	var udto UserDTO
	err := copier.Copy(&udto, &u)
	if err != nil {
		return nil, err
	}
	return &udto, nil
}

func (us *userService) Add() error {
	return serr.NewServiceError("my name is error")
}

func (us *userService) Update() {

}
