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
			models.Endpoint{Path: "register", Method: http.HTTP_METHOD_PUT, Listener: services.RegisterHttp},
		},
	)

}
