package handler

import (
	"fmt"
	"net/http"

	"github.com/costa38r/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("id cannot be empyt: %s", id))
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "Not found")
		return
	}

	sendSuccess(c, "show opening", opening)

}
