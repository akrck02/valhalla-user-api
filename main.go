package main

import (
	apicommon "github.com/akrck02/valhalla-api-common"
	"github.com/akrck02/valhalla-api-common/configuration"
	"github.com/akrck02/valhalla-api-common/models"
	"github.com/akrck02/valhalla-core-sdk/http"
	"github.com/akrck02/valhalla-user-api/services"
)

// main function
func main() {

	config := configuration.LoadConfiguration(".env")

	apicommon.Start(
		config,
		[]models.Endpoint{

			// User endpoints
			{Path: "register", Method: http.HTTP_METHOD_PUT, Listener: services.RegisterHttp},
			{Path: "login", Method: http.HTTP_METHOD_POST, Listener: services.LoginHttp},
			{Path: "login/auth", Method: http.HTTP_METHOD_POST, Listener: services.LoginAuthHttp},
			{Path: "edit", Method: http.HTTP_METHOD_PATCH, Listener: services.EditUserHttp},
			{Path: "edit/email", Method: http.HTTP_METHOD_PATCH, Listener: services.EditUserEmailHttp},
			{Path: "edit/profilepic", Method: http.HTTP_METHOD_PATCH, Listener: services.EditUserProfilePictureHttp},
			{Path: "delete", Method: http.HTTP_METHOD_DELETE, Listener: services.DeleteUserHttp},
			{Path: "get", Method: http.HTTP_METHOD_GET, Listener: services.GetUserHttp},
			{Path: "validate", Method: http.HTTP_METHOD_GET, Listener: services.ValidateUserHttp},

			// Role endpoints
			{Path: "role/create", Method: http.HTTP_METHOD_PUT, Listener: services.CreateRoleHttp},
			{Path: "role/delete", Method: http.HTTP_METHOD_DELETE, Listener: services.DeleteRoleHttp},
			{Path: "role/get", Method: http.HTTP_METHOD_GET, Listener: services.GetRoleHttp},
			{Path: "role/edit", Method: http.HTTP_METHOD_PATCH, Listener: services.EditRoleHttp},

			// Team endpoints
			{Path: "team/create", Method: http.HTTP_METHOD_PUT, Listener: services.CreateTeamHttp},
			{Path: "team/delete", Method: http.HTTP_METHOD_DELETE, Listener: services.DeleteTeamHttp},
			{Path: "team/get", Method: http.HTTP_METHOD_GET, Listener: services.GetTeamHttp},
			{Path: "team/edit", Method: http.HTTP_METHOD_PATCH, Listener: services.EditTeamHttp},
			{Path: "team/edit/owner", Method: http.HTTP_METHOD_PUT, Listener: services.EditTeamOwnerHttp},
			{Path: "team/add/member", Method: http.HTTP_METHOD_PUT, Listener: services.AddMemberHttp},
			{Path: "team/remove/member", Method: http.HTTP_METHOD_DELETE, Listener: services.RemoveMemberHttp},
		},
	)

}
