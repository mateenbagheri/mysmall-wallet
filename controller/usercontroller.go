package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mateenbagheri/mysmall-wallet/service"
)

func GetUserBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		id_mask := c.Param("UserID")

		if id_mask == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "could not find user_id in parameters",
				"error":   "user_id == \"\"",
			})
			return
		}

		id, err := strconv.Atoi(id_mask)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "user id needs to be an integer.",
				"error":   string(err.Error()),
			})
			return
		}

		balance, err := service.SelectUserBalanceByID(id)

		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "could not select this users balance from database",
				"error":   string(err.Error()),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, balance)
	}
}
