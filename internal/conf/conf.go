package conf

import (
	"fastgo/internal/data"
	"fastgo/internal/service/user"
	"fastgo/pkg/di"
	"github.com/eininst/rlock"
)

func Inject() {
	//inject resources
	rcli := data.NewRedisClient()
	//db := data.NewDB()
	lock := rlock.New(rcli)

	di.Inject(rcli)
	//di.Inject(db)
	di.Inject(lock)

	//inject services
	di.Inject(user.NewUserService())
}
