package services

import (
	"github.com/akrck02/valhalla-api-common/models"
	"github.com/akrck02/valhalla-core-sdk/http"
)

var CreateTeamEndpoint = models.Endpoint{
	Path:     "team/create",
	Method:   http.HTTP_METHOD_PUT,
	Listener: CreateTeam,
	Checks:   CreateTeamCheck,
	Database: true,
	Secured:  true,
}

var DeleteTeamEndpoint = models.Endpoint{
	Path:     "team/delete",
	Method:   http.HTTP_METHOD_DELETE,
	Listener: DeleteTeam,
	Checks:   DeleteTeamCheck,
	Database: true,
	Secured:  true,
}

var GetTeamEndpoint = models.Endpoint{
	Path:     "team/get",
	Method:   http.HTTP_METHOD_GET,
	Listener: GetTeam,
	Checks:   GetTeamCheck,
	Database: true,
	Secured:  true,
}

var EditTeamEndpoint = models.Endpoint{
	Path:     "team/edit",
	Method:   http.HTTP_METHOD_POST,
	Listener: EditTeam,
	Checks:   EditTeamCheck,
	Database: true,
	Secured:  true,
}

var EditTeamOwnerEndpoint = models.Endpoint{
	Path:     "team/edit/owner",
	Method:   http.HTTP_METHOD_PUT,
	Listener: EditTeamOwner,
	Checks:   EditTeamOwnerCheck,
	Database: true,
	Secured:  true,
}

var AddMemberEndpoint = models.Endpoint{
	Path:     "team/member/add",
	Method:   http.HTTP_METHOD_PUT,
	Listener: AddMember,
	Checks:   AddMemberCheck,
	Database: true,
	Secured:  true,
}

var RemoveMemberEndpoint = models.Endpoint{
	Path:     "team/member/remove",
	Method:   http.HTTP_METHOD_DELETE,
	Listener: RemoveMember,
	Checks:   RemoveMemberCheck,
	Database: true,
	Secured:  true,
}
