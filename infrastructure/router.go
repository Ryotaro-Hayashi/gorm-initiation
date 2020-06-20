package infrastructure

import (
	"github.com/gin-gonic/gin"
	"gorm-initiation/interface/controllers"
)

// init関数内のルーターをserver.goでimportしてもundefinedになってしまうのでinit関数外でルーターを初期化
var Router *gin.Engine

func init() {
	router := gin.Default()

	userController := controllers.NewUserController(NewSqlHandler())

	// 第2引数は関数でないといけない
	router.GET("/users/get", func(c *gin.Context) { userController.Index(c) })

	Router = router
}
