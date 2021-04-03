package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Url     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func Load() []Route {
	routes := userRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load(){
		r.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
	
	return r
}
