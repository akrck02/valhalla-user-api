package services

import (
	permissiondal "github.com/akrck02/valhalla-core-dal/services/permission"
	userdal "github.com/akrck02/valhalla-core-dal/services/user"
	"github.com/akrck02/valhalla-core-sdk/http"
	devicemodels "github.com/akrck02/valhalla-core-sdk/models/device"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"
)

func Register(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	user := context.Request.Body.(*usersmodels.User)
	error := userdal.Register(context.Database.Client, user)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "User created"},
	}, nil
}

func Login(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	user := context.Request.Body.(*usersmodels.User)
	token, error := userdal.Login(context.Database.Client, user, context.Request.IP, context.Request.UserAgent)

	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code: http.HTTP_STATUS_OK,
		Response: devicemodels.Device{
			User:      user.Email,
			Token:     token,
			Address:   context.Request.IP,
			UserAgent: context.Request.UserAgent,
		},
	}, nil
}

func LoginAuth(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	auth := context.Request.Body.(*usersmodels.AuthLogin)

	error := userdal.LoginAuth(context.Database.Client, auth, context.Request.IP, context.Request.UserAgent)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "User logged in"},
	}, nil

}

func EditUser(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	// check if request is made by humans ;)
	if context.Launcher.LauncherType != systemmodels.USER {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	// get if request user can edit the user
	userToEdit := context.Request.Body.(*usersmodels.User)
	canEdit := permissiondal.CanEditUser(context.Request.User, userToEdit)
	if !canEdit {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Cannot edit user",
		}
	}

	updateErr := userdal.EditUser(context.Database.Client, userToEdit)
	if updateErr != nil {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "User updated"},
	}, nil
}

func EditUserEmail(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	// check if request is made by humans ;)
	if context.Launcher.LauncherType != systemmodels.USER {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	emailChangeRequest := context.Request.Body.(*userdal.EmailChangeRequest)

	// get if request user can edit the user
	canEdit := permissiondal.CanEditUser(context.Request.User, &usersmodels.User{Email: emailChangeRequest.Email})
	if !canEdit {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
	}

	changeErr := userdal.EditUserEmail(context.Database.Client, emailChangeRequest)
	if changeErr != nil {
		return nil, changeErr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Email updated"},
	}, nil
}

func DeleteUser(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	// check if request is made by humans ;)
	if context.Launcher.LauncherType != systemmodels.USER {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	// get if request user can delete the user
	user := context.Request.Body.(*usersmodels.User)
	requestingUser, err := userdal.GetUser(context.Database.Client, &usersmodels.User{ID: context.Launcher.Id}, false)

	if err != nil {
		return nil, err
	}

	canDelete := permissiondal.CanEditUser(requestingUser, user)
	if !canDelete {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot delete user",
		}
	}

	// delete the user
	deleteErr := userdal.DeleteUser(context.Database.Client, user)
	if deleteErr != nil {
		return nil, deleteErr
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "User deleted"},
	}, nil
}

func GetUser(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	requestingUser, err := userdal.GetUser(context.Database.Client, &usersmodels.User{ID: context.Launcher.Id}, false)

	if err != nil {
		return nil, err
	}

	user := context.Request.Body.(*usersmodels.User)

	// get if request user can see the user
	canSee := permissiondal.CanSeeUser(requestingUser, user)
	if !canSee {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot see the user",
		}
	}

	foundUser, error := userdal.GetUser(context.Database.Client, user, true)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: foundUser,
	}, nil

}

func EditUserProfilePicture(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	user := context.Request.Body.(*usersmodels.User)
	bytes := context.Request.Body.([]byte)

	// get if request user can delete the user
	requestingUser, err := userdal.GetUser(context.Database.Client, &usersmodels.User{ID: context.Launcher.Id}, false)

	if err != nil {
		return nil, err
	}

	canEdit := permissiondal.CanEditUser(requestingUser, user)
	if !canEdit {
		return nil, &systemmodels.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
	}

	// Upload image
	error := userdal.EditUserProfilePicture(context.Database.Client, user, bytes)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "Profile picture updated"},
	}, nil

}

func ValidateUser(context *systemmodels.ValhallaContext) (*systemmodels.Response, *systemmodels.Error) {

	// Get code from url GET parameter
	code := context.Request.Body.(string)
	error := userdal.ValidateUser(context.Database.Client, code)
	if error != nil {
		return nil, error
	}

	return &systemmodels.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: systemmodels.Message{Message: "User validated"},
	}, nil
}
