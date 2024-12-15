package handler

import (
	"net/http"

	"github.com/costa38r/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpenings(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "listing", openings)
}
