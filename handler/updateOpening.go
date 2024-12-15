package handler

import (
	"net/http"

	"github.com/costa38r/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpadateOpeningHandler(c *gin.Context) {
	request := UpadateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err)
		sendError(c, http.StatusBadRequest, err.Error())
	}

	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
		return
	}

	//update opening

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salaray > 0 {
		opening.Salary = request.Salaray
	}

	//save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("error updating opening %v", err)
		sendError(c, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(c, "update opening", opening)

}
