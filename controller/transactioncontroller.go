package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mateenbagheri/mysmall-wallet/models"
	"github.com/mateenbagheri/mysmall-wallet/service"
)

func AddTransaction(c *gin.Context) {
	var body models.Transaction

	err := c.BindJSON(&body)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "could not match body with structure required",
			"error":   err.Error(),
		})
		return
	}

	tmpCreateDate := time.Now()
	body.TransactionDate = tmpCreateDate.Format("2006-01-02")

	reference, err := service.InsertTransaction(body)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "could not insert data into database",
			"error":   string(err.Error()),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, reference)
}
