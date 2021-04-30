package controllers

import (
	"api/database"
	"api/models"
	"api/repository"
	"api/repository/crud"
	"api/responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewPostSRepositoryCRUD(db)

	func(postRepository repository.PostRepository) {
		posts, err := postRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, posts)

	}(repo)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {

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

	func(postRepository repository.PostRepository) {
		post, err = postRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Locations", fmt.Sprintf("%s%s%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusCreated, post)

	}(repo)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewPostSRepositoryCRUD(db)

	func(postRepository repository.PostRepository) {
		post, err := postRepository.FindById(uint64(pid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

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

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewPostSRepositoryCRUD(db)
	func(postRepository repository.PostRepository) {
		rows, err := postRepository.Update(pid, post)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusCreated, rows)

	}(repo)

}

func DeletePost(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	repo := crud.NewPostSRepositoryCRUD(db)

	func(postRepository repository.PostRepository) {
		_, err := postRepository.Delete(pid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", pid))
		responses.JSON(w, http.StatusNoContent, "")

	}(repo)
}
