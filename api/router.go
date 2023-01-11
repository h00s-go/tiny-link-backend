package api

import "github.com/h00s-go/tiny-link-backend/api/controllers"

func (api *API) SetRoutes() {
	healthController := controllers.NewHealthController(api.services)
	linksController := controllers.NewLinksController(api.services)

	api.server.Get("/api/v1/health", healthController.GetHealthHandler)
	api.server.Get("/api/v1/links/:id", linksController.GetLinkHandler)
}
