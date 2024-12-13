package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpenings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
