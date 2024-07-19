package services

import (
	roledal "github.com/akrck02/valhalla-core-dal/services/role"
	"github.com/akrck02/valhalla-core-sdk/http"
	rolemodels "github.com/akrck02/valhalla-core-sdk/models/role"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	"github.com/akrck02/valhalla-core-sdk/valerror"
)

func CreateRole(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	var role rolemodels.Role
	var error = roledal.CreateRole(role)

	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: "Role created",
	}, nil

}

func DeleteRole(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {
	return nil, &systemmodels.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

func EditRole(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {
	return nil, &systemmodels.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}

func GetRole(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {
	return nil, &systemmodels.Error{
		Status:  http.HTTP_STATUS_NOT_IMPLEMENTED,
		Error:   valerror.NOT_IMPLEMENTED,
		Message: "Not implemented",
	}
}
