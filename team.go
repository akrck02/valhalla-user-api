package services

import (
	"github.com/akrck02/valhalla-api-common/http"
	"github.com/akrck02/valhalla-core-dal/database"
	"github.com/akrck02/valhalla-core-sdk/error"
	"github.com/akrck02/valhalla-core-sdk/models"
	"github.com/gin-gonic/gin"
)

// Create team HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func CreateTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var team *models.Team = &models.Team{}

	err := c.ShouldatabaseindJSON(team)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = CreateTeam(conn, client, team)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team created"},
	}, nil

}

// Edit team HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func EditTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var params *models.Team = &models.Team{}
	err := c.ShouldatabaseindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = EditTeam(conn, client, params)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team changed"},
	}, nil
}

// Edit team owner HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func EditTeamOwnerHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var params *models.Team = &models.Team{}

	err := c.ShouldatabaseindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = EditTeamOwner(conn, client, params)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team owner edited"},
	}, nil
}

// Delete team HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func DeleteTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var params *models.Team = &models.Team{}
	err := c.ShouldatabaseindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = DeleteTeam(conn, client, params)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team deleted"},
	}, nil
}

// Get team HTTP API endpoint
//
// [param] request | models.Request: request
//
// [return] *models.Response: response | *models.Error: error
func GetTeamHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var params models.Team = models.Team{}
	params.ID = c.Query("id")

	if params.ID == "" {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Team ID is required",
		}
	}

	team, error := GetTeam(conn, client, &params)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team found", "team": team},
	}, nil
}

// Add user to team HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func AddMemberHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)
	var params *MemberChangeRequest = &MemberChangeRequest{}

	err := c.ShouldatabaseindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var addMemberErr = AddMember(conn, client, params)
	if err != nil {
		return nil, addMemberErr
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member added"},
	}, nil
}

// Remove user from team HTTP API endpoint
//
// [param] c | *gin.Context: context
//
// [return] *models.Response: response | *models.Error: error
func RemoveMemberHttp(c *gin.Context) (*models.Response, *models.Error) {

	var client = database.CreateClient()
	var conn = database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var params *MemberChangeRequest = &MemberChangeRequest{}
	err := c.ShouldatabaseindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var removeMemberErr = RemoveMember(conn, client, params)
	if err != nil {
		return nil, removeMemberErr
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member removed"},
	}, nil
}
