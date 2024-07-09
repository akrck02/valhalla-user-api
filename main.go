package main

import (
	apicommon "github.com/akrck02/valhalla-api-common"
	"github.com/akrck02/valhalla-api-common/configuration"
	"github.com/akrck02/valhalla-api-common/models"
	databaseConfig "github.com/akrck02/valhalla-core-dal/configuration"
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
			services.RegisterEndpoint,
			services.LoginEndpoint,
			services.LoginAuthEndpoint,
			services.EditUserEndpoint,
			services.EditUserPasswordEndpoint,
			services.EditUserProfilePictureEndpoint,
			services.DeleteUserEndpoint,
			services.GetUserEndpoint,
			services.ValidateUserEndpoint,

			// Role endpoints
			services.CreateRoleEndpoint,
			services.DeleteRoleEndpoint,
			services.GetRoleEndpoint,
			services.EditRoleEndpoint,

			// Team endpoints
			services.CreateTeamEndpoint,
			services.DeleteTeamEndpoint,
			services.GetTeamEndpoint,
			services.EditTeamEndpoint,
			services.AddMemberEndpoint,
			services.RemoveMemberEndpoint,
		},
	)
}
