package routes

import (
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	TransactionRoute(r)
	UserRoute(r)
	return r
}
