package api

import "github.com/h00s-go/tiny-link-backend/api/controllers"

func (api *API) SetRoutes() {
	healthController := controllers.NewHealthController(api.services)

	api.server.GET("/api/v1/health", healthController.GetHealthHandler)
}
