package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/model"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/service"
	"github.com/gorilla/mux"
)

// AuthorList
func (c *Controller) ListAuthor(w http.ResponseWriter, req *http.Request) {

	results, err := service.GetAuthorsWithBooks(c.DB)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}

// AuthorGet
func (c *Controller) GetAuthor(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	fmt.Println(id)

	results, err := service.GetByIDWithBooks(c.DB, id)

	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}

// AuthorDelete
func (c *Controller) DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	err := service.DeleteAuthorByID(c.DB, id)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)

}

// AuthorCreate
func (c *Controller) CreateAuthor(w http.ResponseWriter, req *http.Request) {

	var author *model.Author

	if err := json.NewDecoder(req.Body).Decode(&author); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	author, err := service.CreateAuthor(c.DB, author)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, author)

}

// UpdateAuthor controller
func (c *Controller) UpdateAuthor(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Author ID")
		return
	}

	author := new(model.Author)

	if err := json.NewDecoder(req.Body).Decode(&author); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err = service.GetAuthorByID(c.DB, id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Author not found")
		return
	}
	defer req.Body.Close()

	author.ID = uint(id)

	author, err = service.UpdateAuthor(c.DB, author)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, author)

}
