package conf

import (
	"fastgo/internal/data/rdb"
	"fastgo/internal/service/user"
	"fastgo/pkg/di"
	"github.com/eininst/rlock"
)

func Inject() {
	rcli := rdb.New()
	di.Inject(rcli)
	//ioc.Provide(db.New())
	di.Inject(rlock.New(rcli))

	//inject services
	di.Inject(user.NewUserService())
}
