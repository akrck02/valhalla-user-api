package services

import (
	server "github.com/akrck02/valhalla-api-common/http"
	"github.com/akrck02/valhalla-core-dal/database"
	permissiondal "github.com/akrck02/valhalla-core-dal/services/permission"
	userdal "github.com/akrck02/valhalla-core-dal/services/user"
	"github.com/akrck02/valhalla-core-sdk/error"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-core-sdk/log"
	"github.com/akrck02/valhalla-core-sdk/models"

	"github.com/gin-gonic/gin"
)

// Register HTTP API endpoint
//
// [param] c | *gin.Context: context
func RegisterHttp(c *gin.Context) (*models.Response, *models.Error) {

	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	error := userdal.Register(conn, client, user)
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
func LoginHttp(c *gin.Context) (*models.Response, *models.Error) {

	request := server.GetRequestMetadata(c)
	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	var user *models.User = &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	ip := request.IP
	address := request.UserAgent
	token, error := userdal.Login(conn, client, user, ip, address)

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
	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	auth := &models.AuthLogin{}
	err := c.ShouldBindJSON(auth)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	auth.AuthToken = request.Authorization

	error := userdal.LoginAuth(conn, client, auth, request.IP, request.UserAgent)
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
	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	userToEdit := &models.User{}
	err := c.ShouldBindJSON(userToEdit)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	// get if request user can edit the user
	canEdit := permissiondal.CanEditUser(request.User, userToEdit)
	if !canEdit {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Cannot edit user",
		}
	}

	updateErr := userdal.EditUser(conn, client, userToEdit)
	if updateErr != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
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
	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	email := &userdal.EmailChangeRequest{}
	err := c.ShouldBindJSON(email)
	if err != nil {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request",
		}
	}

	// get if request user can edit the user
	canEdit := permissiondal.CanEditUser(request.User, &models.User{Email: email.Email})
	if !canEdit {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
	}

	changeErr := userdal.EditUserEmail(conn, client, email)
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
	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	user := &models.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		return nil, &models.Error{
			Status: http.HTTP_STATUS_BAD_REQUEST,
			Error:  error.INVALID_REQUEST,
		}
	}

	// get if request user can delete the user
	canDelete := permissiondal.CanEditUser(request.User, user)
	if !canDelete {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot delete user",
		}
	}

	deleteErr := userdal.DeleteUser(conn, client, user)
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
	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	// Get code from url GET parameter
	id := c.Query("id")
	if id == "" {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_BAD_REQUEST,
			Error:   error.INVALID_REQUEST,
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
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot see the user",
		}
	}

	foundUser, error := userdal.GetUser(conn, client, user, true)
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
			Error:   error.INVALID_REQUEST,
			Message: "Invalid request body",
		}
	}

	// get if request user can delete the user
	canEdit := permissiondal.CanEditUser(request.User, user)
	if !canEdit {
		return nil, &models.Error{
			Status:  http.HTTP_STATUS_FORBIDDEN,
			Error:   error.ACCESS_DENIED,
			Message: "Access denied: Cannot edit user",
		}
	}

	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	// Upload image
	error := userdal.EditUserProfilePicture(conn, client, user, bytes)
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

	client := database.CreateClient()
	conn := database.Connect(*client)
	defer database.Disconnect(*client, conn)

	// Get code from url GET parameter
	code := c.Query("code")
	log.Info("Query code: " + code)

	error := userdal.ValidateUser(conn, client, code)
	if error != nil {
		return nil, error
	}

	return &models.Response{
		Code:     http.HTTP_STATUS_OK,
		Response: gin.H{"message": "User validated"},
	}, nil
}
