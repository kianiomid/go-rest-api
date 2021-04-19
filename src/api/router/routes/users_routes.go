package routes

import (
	"api/controllers"
	"net/http"
)

var userRoutes = []Route{
	Route{
		Url:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		Url:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		Url:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		Url:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
	Route{
		Url:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
}
