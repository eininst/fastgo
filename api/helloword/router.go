package helloword

import (
	v1 "github.com/eininst/fastgo/api/helloword/v1"
	"github.com/gin-gonic/gin"
)

func Install(r *gin.Engine) {
	r.GET("/hello", v1.HelloWorld)
}
