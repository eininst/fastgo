package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(c *gin.Context) {
	if 1 == 1 {
		panic("weew")
	}
	c.JSON(http.StatusOK, gin.H{"name": "hello wolrd!"})
}
