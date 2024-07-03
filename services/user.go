package services

import (
	server "github.com/akrck02/valhalla-api-common/http"
	permissiondal "github.com/akrck02/valhalla-core-dal/services/permission"
	userdal "github.com/akrck02/valhalla-core-dal/services/user"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-core-sdk/log"
	"github.com/akrck02/valhalla-core-sdk/models"
	systemmodels "github.com/akrck02/valhalla-core-sdk/models/system"
	usersmodels "github.com/akrck02/valhalla-core-sdk/models/users"
	"github.com/akrck02/valhalla-core-sdk/valerror"

	"github.com/gin-gonic/gin"
)

func Register(context systemmodels.ValhallaContext) (*models.Response, *models.Error) {

	user := context.(systemmodels.Request).user
	error := userdal.Register(user)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User created"},
	}, nil
}

// Login HTTP API endpoint
//
// [param] c | *gin.Context: context
func LoginHttp(context systemmodels.ValhallaContext, gin *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(gin)
	var user *usersmodels.User = &usersmodels.User{}
	err := gin.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	ip := request.IP
	address := request.UserAgent
	token, error := userdal.Login(user, ip, address)

	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code: http.HTTP_STATUS_OK,
		Response: gin.H{
			"auth":  token,
			"email": user.Email,
		},
	}, nil
}

// Login auth HTTP API endpoint
//
// [param] c | *gin.Context: context
func LoginAuthHttp(c *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(c)
	auth := &usersmodels.AuthLogin{
		Email:     request.User.Email,
		AuthToken: request.Authorization,
	}

	auth.AuthToken = request.Authorization
	error := userdal.LoginAuth(auth, request.IP, request.UserAgent)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User logged in"},
	}, nil

}

// Edit user HTTP API endpoint
//
// [param] c | *gin.Context: context
func EditUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(c)
	userToEdit := &usersmodels.User{}
	err := c.ShouldBindJSON(userToEdit)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	// get if request user can edit the user
	canEdit := permissiondal.CanEditUser(request.User, userToEdit)
	if !canEdit {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Cannot edit user",
		}
	}

	updateErr := userdal.EditUser(userToEdit)
	if updateErr != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User updated"},
	}, nil
}

// Change email HTTP API endpoint
//
// [param] c | *gin.Context: context
func EditUserEmailHttp(c *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(c)

	email := &userdal.EmailChangeRequest{}
	err := c.ShouldBindJSON(email)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	// get if request user can edit the user
	canEdit := permissiondal.CanEditUser(request.User, &models.User{Email: email.Email})
	if !canEdit {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
	}

	changeErr := userdal.EditUserEmail(email)
	if changeErr != nil {
		return nil, changeErr
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Email changed"},
	}, nil
}

// Delete user HTTP API endpoint
//
// [param] c | *gin.Context: context
func DeleteUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(c)

	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Status: http.HTTP_STATUS_BAD_REQUEST,
			Error:  valerror.INVALID_REQUEST,
		}
	}

	// get if request user can delete the user
	canDelete := permissiondal.CanEditUser(request.User, user)
	if !canDelete {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot delete user",
		}
	}

	deleteErr := userdal.DeleteUser(user)
	if deleteErr != nil {
		return nil, deleteErr
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User deleted"},
	}, nil
}

// Get user HTTP API endpoint
//
// [param] c | *gin.Context: context
func GetUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(c)

	// Get code from url GET parameter
	id := c.Query("id")
	if id == "" {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Id cannot be empty",
		}
	}

	var user *models.User = &models.User{
		Email: id,
	}

	// get if request user can see the user
	canSee := permissiondal.CanSeeUser(request.User, user)
	if !canSee {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot see the user",
		}
	}

	foundUser, error := userdal.GetUser(user, true)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User found", "user": foundUser},
	}, nil

}

// Edit user profile picture HTTP API endpoint
//
// [param] c | *gin.Context: context
func EditUserProfilePictureHttp(c *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(c)

	// Get user
	user := &models.User{
		Email: c.Query("email"),
	}

	// Get image as bytes
	bytes, err := server.MultipartToBytes(c, "ProfilePicture")
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   valerror.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	// get if request user can delete the user
	canEdit := permissiondal.CanEditUser(request.User, user)
	if !canEdit {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   valerror.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
	}

	// Upload image
	error := userdal.EditUserProfilePicture(user, bytes)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "Profile picture updated"},
	}, nil

}

// Validate user account HTTP API endpoint
//
// [param] c | *gin.Context: context
func ValidateUserHttp(c *gin.Context) (*models.Response, *models.Error) {

	// Get code from url GET parameter
	code := c.Query("code")
	log.Info("Query code: " + code)

	error := userdal.ValidateUser(code)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User validated"},
	}, nil
}
