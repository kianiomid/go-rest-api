package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects/go-rest-api/src/api/models"
	"projects/go-rest-api/src/api/repository"
	"projects/go-rest-api/src/api/repository/crud"
	"projects/go-rest-api/src/api/responses"
	"projects/go-rest-api/src/database"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of users"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewUsersRepositoryCRUD(db)

	func (usersRepository repository.UserRepository){
		user, err = usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Locations", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)

	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("An user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
