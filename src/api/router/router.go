package router

import (
	"github.com/gorilla/mux"
	"projects/go-rest-api/src/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
