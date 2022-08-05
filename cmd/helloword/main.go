package main

import (
	"fmt"
	"github.com/eininst/fastgo/api/helloword"
	"github.com/eininst/fastgo/configs"
	"github.com/eininst/fastgo/tools/grace"
	"github.com/eininst/fastgo/tools/serr"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	rand.Seed(time.Now().UnixNano())

	configs.Setup("configs/config.yml")
	//db.Setup()
	//rdb.Setup()
}
func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.CustomRecovery(serr.ErrorRecovery))

	helloword.Install(r)
	addr := fmt.Sprintf(":%v", "8080")

	grace.Run("./main.go", &http.Server{
		Addr:    addr,
		Handler: r,
	}, time.Second*10)
}
