package services

import (
	"github.com/akrck02/valhalla-core-sdk/http"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	teammodels "github.com/akrck02/valhalla-core-sdk/models/team"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func CreateTeamCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	team := &teammodels.Team{}
	err := gin.ShouldBindJSON(team)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	context.Request.Body = team
	return nil

}

func EditTeamCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	team := &teammodels.Team{}
	err := gin.ShouldBindJSON(team)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	context.Request.Body = team
	return nil

}

func EditTeamOwnerCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	team := &teammodels.Team{}
	err := gin.ShouldBindJSON(team)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	context.Request.Body = team
	return nil

}

func DeleteTeamCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	team := &teammodels.Team{}
	err := gin.ShouldBindJSON(team)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	context.Request.Body = team
	return nil

}

func GetTeamCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	team := &teammodels.Team{}
	team.ID = gin.Query("id")

	if team.ID == "" {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Team ID is required",
		}
	}

	return nil

}

func AddMemberCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	memberChangeRequest := &teammodels.MemberChangeRequest{}
	err := gin.ShouldBindJSON(memberChangeRequest)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	context.Request.Body = memberChangeRequest
	return nil

}

func RemoveMemberCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	memberChangeRequest := &teammodels.MemberChangeRequest{}
	err := gin.ShouldBindJSON(memberChangeRequest)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	context.Request.Body = memberChangeRequest
	return nil

}
