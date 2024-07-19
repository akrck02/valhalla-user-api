package services

import (
	"github.com/akrck02/valhalla-api-common/server"
	userdal "github.com/akrck02/valhalla-core-dal/services/user"
	"github.com/akrck02/valhalla-core-sdk/http"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"
	"github.com/gin-gonic/gin"
)

func RegisterCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	user := &usersmodels.User{}
	err := gin.ShouldBindJSON(user)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	context.Request.Body = user
	return nil
}

func LoginCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	var user *usersmodels.User = &usersmodels.User{}
	err := gin.ShouldBindJSON(user)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	context.Request.Body = user
	return nil
}

func LoginAuthCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	request := server.GetRequestMetadata(gin)
	auth := &usersmodels.AuthLogin{
		Email:     request.User.Email,
		AuthToken: request.Authorization,
	}

	auth.AuthToken = request.Authorization
	context.Request.Body = auth
	return nil

}

func EditUserCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	userToEdit := &usersmodels.User{}
	err := gin.ShouldBindJSON(userToEdit)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	context.Request.Body = userToEdit
	return nil
}

func EditUserEmailCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	email := &userdal.EmailChangeRequest{}
	err := gin.ShouldBindJSON(email)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	context.Request.Body = email
	return nil

}

func DeleteUserCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	user := &usersmodels.User{}
	err := gin.ShouldBindJSON(user)
	if err != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	context.Request.Body = user
	return nil
}

func GetUserCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	// get code from url GET parameter
	id := gin.Query("id")
	if id == "" {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Id cannot be empty",
		}
	}

	// get user from database
	user, err := userdal.GetUser(context.Database.Client, &usersmodels.User{ID: id}, false)
	if err != nil {
		return err
	}

	context.Request.Body = user
	return nil
}

func EditUserProfilePictureCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	// Get user
	user, err := userdal.GetUser(context.Database.Client, &usersmodels.User{ID: gin.Query("id")}, false)

	if err != nil {
		return err
	}

	// Get image as bytes
	bytes, merr := server.MultipartToBytes(gin, "ProfilePicture")
	if merr != nil {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	context.Request.Body = user
	context.Request.Files = bytes
	return nil
}

func ValidateUserCheck(context *systemmodels.ValhallaContext, gin *gin.Context) *systemmodels.Error {

	// Get code from url GET parameter
	code := gin.Query("code")
	if code == "" {
		return &systemmodels.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Code cannot be empty",
		}
	}

	context.Request.Body = code
	return nil
}
