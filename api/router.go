package api

import (
	"github.com/labstack/echo/v4"
)

func (api *API) NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/api/v1/health", api.GetHealthHandler)

	return e
}
