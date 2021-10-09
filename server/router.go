package server

import (
	"net/http"

	"crud_api/handler"
)

func Router() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/ping", Pong)

	userAPI := handler.UserAPI{}
	router.HandleFunc("/users", userAPI.CreateUser)
	router.HandleFunc("/users/id", userAPI.GetUserFromID)

	postAPI := handler.PostAPI{User: userAPI}
	router.HandleFunc("/posts", postAPI.CreatePost)
	router.HandleFunc("/posts/id", postAPI.GetPostUsingID)
	router.HandleFunc("/posts/users/id", postAPI.GetAllPostsOfUser)

	return router
}
