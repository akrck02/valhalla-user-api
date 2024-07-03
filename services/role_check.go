package services

import (
	"github.com/akrck02/valhalla-core-sdk/http"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func CreateRoleCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	var params usersmodels.User
	err := gin.ShouldBindJSON(&params)

	if err != nil {
		return nil, &usersmodels.Error{
			Status:  http.HTTP_STATUS_NOT_ACCEPTABLE,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return nil
}
