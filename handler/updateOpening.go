package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpadateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
