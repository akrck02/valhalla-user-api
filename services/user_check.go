package services

import (
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-core-sdk/models"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func RegisterCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	user := &usersmodels.User{}
	err := gin.ShouldBindJSON(user)
	if err != nil {
		return &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return nil
}
