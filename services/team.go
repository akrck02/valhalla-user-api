package services

import (
	teamdal "github.com/akrck02/valhalla-core-dal/services/team"
	"github.com/akrck02/valhalla-core-sdk/http"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	teammodels "github.com/akrck02/valhalla-core-sdk/models/team"
)

func CreateTeam(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	team := context.Request.Body.(*teammodels.Team)

	var error = teamdal.CreateTeam(team)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Team created"},
	}, nil

}

func EditTeam(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	team := context.Request.Body.(*teammodels.Team)
	error := teamdal.EditTeam(team)

	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Team changed"},
	}, nil
}

func EditTeamOwner(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	team := context.Request.Body.(*teammodels.Team)
	error := teamdal.EditTeamOwner(team)

	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Team owner edited"},
	}, nil
}

func DeleteTeam(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	team := context.Request.Body.(*teammodels.Team)
	var error = teamdal.DeleteTeam(team)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Team deleted"},
	}, nil
}

func GetTeam(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	team := context.Request.Body.(*teammodels.Team)
	team, error := teamdal.GetTeam(team)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: team,
	}, nil
}

func AddMember(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	memberChangeRequest := context.Request.Body.(*teamdal.MemberChangeRequest)
	var addMemberErr = teamdal.AddMember(memberChangeRequest)
	if addMemberErr != nil {
		return nil, addMemberErr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Member added"},
	}, nil
}

func RemoveMember(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	memberChangeRequest := context.Request.Body.(*teamdal.MemberChangeRequest)
	removeMemberErr := teamdal.RemoveMember(memberChangeRequest)
	if removeMemberErr != nil {
		return nil, removeMemberErr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Member removed"},
	}, nil
}
