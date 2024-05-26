package services

import (
	roledal "github.com/akrck02/valhalla-core-dal/services/role"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-core-sdk/models"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

// CreateRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func CreateRoleHttp(c *gin.Context) (*models.Response, *models.Error) {

	var params models.User
	err := c.ShouldBindJSON(&params)

	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_NOT_ACCEPTABLE,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	var role models.Role
	var error = roledal.CreateRole(role)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "User created",
	}, nil

}

// DeleteRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func DeleteRoleHttp(c *gin.Context) (*models.Response, *models.Error) {
	return nil, &models.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

// EditRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func EditRoleHttp(c *gin.Context) (*models.Response, *models.Error) {
	return nil, &models.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

// GetRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func GetRoleHttp(c *gin.Context) (*models.Response, *models.Error) {
	return nil, &models.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}
