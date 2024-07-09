package services

import (
	"github.com/akrck02/valhalla-api-common/models"
	"github.com/akrck02/valhalla-core-sdk/http"
)

var RegisterEndpoint = models.Endpoint{
	Path:     "register",
	Method:   http.HTTP_METHOD_PUT,
	Listener: Register,
	Checks:   RegisterCheck,
	Database: true,
}

var LoginEndpoint = models.Endpoint{
	Path:     "login",
	Method:   http.HTTP_METHOD_POST,
	Listener: Login,
	Checks:   LoginCheck,
	Database: true,
}

var LoginAuthEndpoint = models.Endpoint{
	Path:     "login/auth",
	Method:   http.HTTP_METHOD_POST,
	Listener: LoginAuth,
	Checks:   LoginAuthCheck,
	Database: true,
	Secured:  true,
}

var EditUserEndpoint = models.Endpoint{
	Path:     "edit",
	Method:   http.HTTP_METHOD_POST,
	Listener: EditUser,
	Checks:   EditUserCheck,
	Database: true,
	Secured:  true,
}

var EditUserPasswordEndpoint = models.Endpoint{
	Path:     "edit/email",
	Method:   http.HTTP_METHOD_POST,
	Listener: EditUserEmail,
	Checks:   EditUserEmailCheck,
	Database: true,
	Secured:  true,
}

var EditUserProfilePictureEndpoint = models.Endpoint{
	Path:     "edit/profilepic",
	Method:   http.HTTP_METHOD_POST,
	Listener: EditUserProfilePicture,
	Checks:   EditUserProfilePictureCheck,
	Database: true,
	Secured:  true,
}

var DeleteUserEndpoint = models.Endpoint{
	Path:     "delete",
	Method:   http.HTTP_METHOD_DELETE,
	Listener: DeleteUser,
	Checks:   DeleteUserCheck,
	Database: true,
	Secured:  true,
}

var GetUserEndpoint = models.Endpoint{
	Path:     "get",
	Method:   http.HTTP_METHOD_GET,
	Listener: GetUser,
	Checks:   GetUserCheck,
	Database: true,
	Secured:  true,
}

var ValidateUserEndpoint = models.Endpoint{
	Path:     "validate",
	Method:   http.HTTP_METHOD_GET,
	Listener: ValidateUser,
	Checks:   ValidateUserCheck,
	Database: true,
	Secured:  true,
}
