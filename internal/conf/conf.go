package conf

import (
	"fastgo/internal/common/inject"
	"fastgo/internal/data"
	"fastgo/internal/service/user"
	"github.com/eininst/rlock"
)

func Provide() {
	//inject resources
	rcli := data.NewRedisClient()
	inject.Provide(rcli)
	inject.Provide(rlock.New(rcli))
	inject.Provide(data.NewRsClient(rcli))

	db := data.NewDB()
	inject.Provide(db)

	//inject services
	inject.Provide(user.NewUserService())
}
