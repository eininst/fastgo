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
	di.Inject(rcli)
	di.Inject(rlock.New(rcli))
	di.Inject(data.NewRsClient(rcli))
	//db := data.NewDB()
	//di.Inject(db)

	//inject services
	di.Inject(user.NewUserService())
}
