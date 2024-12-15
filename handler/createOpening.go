package handler

import (
	"net/http"

	"github.com/costa38r/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(c *gin.Context) {

	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salaray,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(c, "createOpening", opening)

}
