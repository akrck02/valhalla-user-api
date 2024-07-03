package services

import (
	roledal "github.com/akrck02/valhalla-core-dal/services/role"
	"github.com/akrck02/valhalla-core-sdk/http"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func CreateRole(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	var role usersmodels.Role
	var error = roledal.CreateRole(role)

	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "Role created",
	}, nil

}

func DeleteRole(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {
	return nil, &usersmodels.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

func EditRole(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {
	return nil, &systemmodels.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

func GetRole(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {
	return nil, &systemmodels.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}
