package services

import (
	models "github.com/akrck02/valhalla-api-common/models"
	"github.com/akrck02/valhalla-core-sdk/http"
)

var CreateRoleEndpoint = models.Endpoint{
	Path:     "role/create",
	Method:   http.HTTP_METHOD_PUT,
	Listener: CreateRole,
	Checks:   CreateRoleCheck,
	Database: true,
	Secured:  true,
}

var DeleteRoleEndpoint = models.Endpoint{
	Path:     "role/delete",
	Method:   http.HTTP_METHOD_DELETE,
	Listener: DeleteRole,
	// Checks:   EmptyCheck,
	Database: true,
	Secured:  true,
}

var GetRoleEndpoint = models.Endpoint{
	Path:     "role/get",
	Method:   http.HTTP_METHOD_GET,
	Listener: GetRole,
	// Checks:   EmptyCheck,
	Database: true,
	Secured:  true,
}

var EditRoleEndpoint = models.Endpoint{
	Path:     "role/edit",
	Method:   http.HTTP_METHOD_POST,
	Listener: EditRole,
	// Checks:   EmptyCheck,
	Database: true,
	Secured:  true,
}
