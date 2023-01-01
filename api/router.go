package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (api *API) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/api/v1/health", api.GetHealthHandler)

	return r
}
