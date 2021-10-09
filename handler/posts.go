package handler

import (
	"net/http"
	"strconv"
	"time"

	"crud_api/contracts"
	"crud_api/database"
	"crud_api/response"
)

type PostAPI struct{
	User UserAPI
}

func (api *PostAPI) CreatePost(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id, err := strconv.ParseInt(query["id"][0], 10, 64)
	if err != nil {
		return
	}
	caption := query["caption"]
	image := query["image"]

	var values []contracts.Posts

	values = append(values, contracts.Posts{
		Id:      id,
		Caption: caption[0],
		Image:   image[0],
		Posted:  time.Now(),
	})

	response.WriterJSON(values, w)
	database.CreatePost(values)
}

func (api *PostAPI) GetPostUsingID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query()["id"][0], 10, 64)
	if err != nil {
		return
	}

	database.GetPostFromID(int(id))
}

func (api *PostAPI) GetAllPostsOfUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query()["id"][0], 10, 64)
	if err != nil {
		return
	}

	database.GetAllPostsOfUser(int(userID))
}
