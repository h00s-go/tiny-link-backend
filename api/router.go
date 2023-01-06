package api

func (api *API) SetRoutes() {
	api.server.GET("/api/v1/health", api.GetHealthHandler)
}
