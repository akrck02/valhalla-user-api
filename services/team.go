package services

import (
	"github.com/akrck02/valhalla-core-dal/database"
	teamdal "github.com/akrck02/valhalla-core-dal/services/team"
	"github.com/akrck02/valhalla-core-sdk/http"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func CreateTeamHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	var team *usersmodels.Team = &usersmodels.Team{}

	err := c.ShouldBindJSON(team)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  systemmodels.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.CreateTeam(team)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team created"},
	}, nil

}

func EditTeamHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	var params *usersmodels.Team = &usersmodels.Team{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.EditTeam(params)

	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team changed"},
	}, nil
}

func EditTeamOwnerHttp(c *gin.Context) (*usersmodels.Response, *systemmodels.Error) {

	var params *usersmodels.Team = &usersmodels.Team{}

	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.EditTeamOwner(params)

	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team owner edited"},
	}, nil
}

func DeleteTeamHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	var params *usersmodels.Team = &usersmodels.Team{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &usersmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var error = teamdal.DeleteTeam(params)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team deleted"},
	}, nil
}

func GetTeamHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	var params usersmodels.Team = usersmodels.Team{}
	params.ID = c.Query("id")

	if params.ID == "" {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Team ID is required",
		}
	}

	team, error := teamdal.GetTeam(&params)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Team found", "team": team},
	}, nil
}

func AddMemberHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	var client = database.Connect()
	defer database.Disconnect(*client)

	var params *teamdal.MemberChangeRequest = &teamdal.MemberChangeRequest{}

	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var addMemberErr = teamdal.AddMember(params)
	if addMemberErr != nil {
		return nil, addMemberErr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member added"},
	}, nil
}

func RemoveMemberHttp(c *gin.Context) (*systemmodels.Response, *systemmodels.Error) {

	var client = database.Connect()
	defer database.Disconnect(*client)

	var params *teamdal.MemberChangeRequest = &teamdal.MemberChangeRequest{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	var removeMemberErr = teamdal.RemoveMember(params)
	if removeMemberErr != nil {
		return nil, removeMemberErr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Member removed"},
	}, nil
}
