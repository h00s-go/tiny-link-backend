package api

import "net/http"

func (api *API) NewRouter() *http.ServeMux {
	return http.NewServeMux()
}
