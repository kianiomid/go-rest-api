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
}
