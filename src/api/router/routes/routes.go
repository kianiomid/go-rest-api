package routes

import (
	"api/middlewares"
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
	routes = append(routes, postRoutes...)
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(
			route.Url,
			middlewares.SetMiddlewareLogger(middlewares.SetMiddlewareJSON(route.Handler)),
			).Methods(route.Method)
	}
	return r
}
