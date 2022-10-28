package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateenbagheri/mysmall-wallet/controller"
)

func UserRoute(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/:UserID/balance", controller.GetUserBalance())
	}
}
