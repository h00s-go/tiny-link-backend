package api

import (
	"github.com/go-chi/chi/v5"
)

func (api *API) NewRouter() *chi.Mux {
	return chi.NewRouter()
}
