package routes

import (
	"api/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		Url:     "/posts",
		Method:  http.MethodGet,
		Handler: controllers.GetPosts,
	},
	{
		Url:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
	},
	{
		Url:     "/posts/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetPost,
	},
	{
		Url:     "/posts/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdatePost,
	},
	{
		Url:     "/posts/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeletePost,
	},
}
