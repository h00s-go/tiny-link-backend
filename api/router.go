package api

import "github.com/h00s-go/tiny-link-backend/api/controllers"

func (api *API) setRoutes() {
	healthController := controllers.NewHealthController(api.services)
	linksController := controllers.NewLinksController(api.services)

	api.server.Get("/api/v1/health", healthController.GetHealthHandler)
	//api.server.Get("/api/v1/links/:id", linksController.GetLinkHandler)
	api.server.Get("/api/v1/links/:short_uri", linksController.GetLinkByShortURIHandler)
	api.server.Post("/api/v1/links", linksController.CreateLinkHandler)
}
