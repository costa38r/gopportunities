package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
