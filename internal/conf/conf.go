package conf

import (
	"fastgo/internal/service/user"
	"fastgo/pkg/di"
)

func Inject() {
	//inject resources
	//rcli := data.NewRedisClient()
	//db := data.NewDB()
	//lock := rlock.New(rcli)

	//di.Inject(rcli)
	//di.Inject(db)
	//di.Inject(lock)

	//inject services
	di.Inject(user.NewUserService())
}
