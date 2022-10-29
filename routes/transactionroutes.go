package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mateenbagheri/mysmall-wallet/controller"
)

func TransactionRoute(router *gin.Engine) {
	transaction := router.Group("/transaction")
	{
		transaction.POST("/", controller.AddTransaction)
	}
}
