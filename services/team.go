package services

import (
	"github.com/akrck02/valhalla-core-dal/database"
	teamdal "github.com/akrck02/valhalla-core-dal/services/team"
	"github.com/akrck02/valhalla-core-sdk/error"
	"github.com/akrck02/valhalla-core-sdk/http"
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

	err := c.ShouldBindJSON(team)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.CreateTeam(conn, client, team)
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
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.EditTeam(conn, client, params)

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

	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.EditTeamOwner(conn, client, params)

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
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.DeleteTeam(conn, client, params)
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

	team, error := teamdal.GetTeam(conn, client, &params)
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
	var params *teamdal.MemberChangeRequest = &teamdal.MemberChangeRequest{}

	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var addMemberErr = teamdal.AddMember(conn, client, params)
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

	var params *teamdal.MemberChangeRequest = &teamdal.MemberChangeRequest{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var removeMemberErr = teamdal.RemoveMember(conn, client, params)
	if err != nil {
		return nil, removeMemberErr
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member removed"},
	}, nil
}
