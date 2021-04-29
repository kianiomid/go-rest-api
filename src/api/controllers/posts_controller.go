package controllers

import (
	"api/database"
	"api/models"
	"api/repository"
	"api/repository/crud"
	"api/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewPostSRepositoryCRUD(db)

	func (postRepository repository.PostRepository){
		posts, err := postRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, posts)

	}(repo)
}

func CreatePost(w http.ResponseWriter, r *http.Request)  {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	err = post.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewPostSRepositoryCRUD(db)

	func (postRepository repository.PostRepository){
		post, err = postRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Locations", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusCreated, post)

	}(repo)
}
