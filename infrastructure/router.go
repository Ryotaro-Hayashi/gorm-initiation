package infrastructure

import (
	"github.com/gin-gonic/gin"
)

// init関数内のルーターをserver.goでimportしてもundefinedになってしまうのでinit関数外でルーターを初期化
var Router *gin.Engine

func init() {
	router := gin.Default()

	router.GET("/user/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success access",
		})
	})

	Router = router
}
