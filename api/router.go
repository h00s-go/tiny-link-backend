package api

import "github.com/h00s-go/tiny-link-backend/api/controllers"

func (api *API) setRoutes() {
	api.server.Get("/api/v1/health", controllers.GetHealthHandler)
	//api.server.Get("/api/v1/links/:id", linksController.GetLinkHandler)
	api.server.Get("/api/v1/links/:short_uri", controllers.GetLinkByShortURIHandler)
	api.server.Post("/api/v1/links", controllers.CreateLinkHandler)
}
