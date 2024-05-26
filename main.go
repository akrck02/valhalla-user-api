package main

import (
	apicommon "github.com/akrck02/valhalla-api-common"
	"github.com/akrck02/valhalla-api-common/configuration"
	"github.com/akrck02/valhalla-api-common/models"
	databaseConfig "github.com/akrck02/valhalla-core-dal/configuration"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-user-api/services"
)

// main function
func main() {

	config := configuration.LoadConfiguration(".env")
	databaseConfig.LoadConfiguration(".env")

	apicommon.Start(
		config,
		[]models.Endpoint{

			// User endpoints
			{Path: "register", Method: http.HTTP_METHOD_PUT, Listener: services.RegisterHttp},
			{Path: "login", Method: http.HTTP_METHOD_POST, Listener: services.LoginHttp},
			{Path: "login/auth", Method: http.HTTP_METHOD_POST, Listener: services.LoginAuthHttp, Secured: true},
			{Path: "edit", Method: http.HTTP_METHOD_POST, Listener: services.EditUserHttp, Secured: true},
			{Path: "edit/email", Method: http.HTTP_METHOD_POST, Listener: services.EditUserEmailHttp, Secured: true},
			{Path: "edit/profilepic", Method: http.HTTP_METHOD_POST, Listener: services.EditUserProfilePictureHttp, Secured: true},
			{Path: "delete", Method: http.HTTP_METHOD_DELETE, Listener: services.DeleteUserHttp, Secured: true},
			{Path: "get", Method: http.HTTP_METHOD_GET, Listener: services.GetUserHttp, Secured: true},
			{Path: "validate", Method: http.HTTP_METHOD_GET, Listener: services.ValidateUserHttp, Secured: true},

			// Role endpoints
			{Path: "role/create", Method: http.HTTP_METHOD_PUT, Listener: services.CreateRoleHttp, Secured: true},
			{Path: "role/delete", Method: http.HTTP_METHOD_DELETE, Listener: services.DeleteRoleHttp, Secured: true},
			{Path: "role/get", Method: http.HTTP_METHOD_GET, Listener: services.GetRoleHttp, Secured: true},
			{Path: "role/edit", Method: http.HTTP_METHOD_POST, Listener: services.EditRoleHttp, Secured: true},

			// Team endpoints
			{Path: "team/create", Method: http.HTTP_METHOD_PUT, Listener: services.CreateTeamHttp, Secured: true},
			{Path: "team/delete", Method: http.HTTP_METHOD_DELETE, Listener: services.DeleteTeamHttp, Secured: true},
			{Path: "team/get", Method: http.HTTP_METHOD_GET, Listener: services.GetTeamHttp, Secured: true},
			{Path: "team/edit", Method: http.HTTP_METHOD_POST, Listener: services.EditTeamHttp, Secured: true},
			{Path: "team/edit/owner", Method: http.HTTP_METHOD_PUT, Listener: services.EditTeamOwnerHttp, Secured: true},
			{Path: "team/add/member", Method: http.HTTP_METHOD_PUT, Listener: services.AddMemberHttp, Secured: true},
			{Path: "team/remove/member", Method: http.HTTP_METHOD_DELETE, Listener: services.RemoveMemberHttp, Secured: true},
		},
	)

}
