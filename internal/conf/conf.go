package conf

import (
	"fastgo/internal/data"
	"fastgo/internal/service/user"
	"github.com/eininst/ninja"
	"github.com/eininst/rlock"
)

func Provide() {
	//inject resources
	rcli := data.NewRedisClient()
	ninja.Provide(rcli)
	ninja.Provide(rlock.New(rcli))
	ninja.Provide(data.NewRsClient(rcli))

	//db := data.NewDB()
	//inject.Provide(db)

	//inject services
	ninja.Provide(user.NewUserService())
}
