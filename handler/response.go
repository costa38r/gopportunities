package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "application/json")

	c.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from: %s successfull", op),
		"data":    data,
	})
}
