package conf

import (
	"fastgo/internal/data/db"
	"fastgo/internal/data/rdb"
	"fastgo/internal/service/user"
	"fastgo/pkg/ioc"
	"github.com/eininst/rlock"
)

func Inject() {
	rcli := rdb.New()
	ioc.Provide(rcli)
	ioc.Provide(db.New())
	ioc.Provide(rlock.New(rcli))

	//inject services
	ioc.Provide(user.NewUserService())
}
