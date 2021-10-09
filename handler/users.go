package handler

import (
	"net/http"
	"strconv"

	"crud_api/contracts"
	"crud_api/database"
	"crud_api/response"
)

type UserAPI struct{}

func (api *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	id, err := strconv.ParseInt(query["id"][0], 10, 64)
	if err != nil {
		return
	}
	name := query["name"]
	email := query["email"]
	password := query["password"]

	var values []contracts.Users

	values = append(values, contracts.Users{
		Id:       id,
		Name:     name[0],
		Email:    email[0],
		Password: password[0],
	})

	response.WriterJSON(values, w)
	database.CreateUser(values)
}

func (api *UserAPI) GetUserFromID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query()["id"][0], 10, 64)
	if err != nil {
		return
	}

	database.GetUserFromID(int(id))
}
