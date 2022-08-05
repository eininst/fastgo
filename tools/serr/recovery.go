package serr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var UNKNOW_ERR = ServiceError{Code: -1, Msg: "未知的系统错误"}

type ServiceError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ErrorRecovery(c *gin.Context, err any) {
	if v, ok := err.(ServiceError); ok {
		if v.Code == 0 {
			v.Code = 400
		}
		c.JSON(http.StatusBadRequest, v)
		c.Abort()
		return
	}
	c.JSON(http.StatusInternalServerError, UNKNOW_ERR)
	c.Abort()
	return
}
