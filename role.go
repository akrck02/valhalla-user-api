package services

import (
	"github.com/akrck02/valhalla-core-dal/database"
	"github.com/akrck02/valhalla-core-sdk/error"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-core-sdk/models"
	"github.com/gin-gonic/gin"
)

// CreateRole HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func CreateRoleHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var params models.User
	err := c.ShouldBindJSON(&params)

	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_NOT_ACCEPTABLE,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	var role models.Role
	var error = CreateRole(conn, client, role)

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
		Error:   error.NOT_IMPLEMENTED,
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
		Error:   error.NOT_IMPLEMENTED,
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
		Error:   error.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}
