package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/service"
	"github.com/gorilla/mux"
)

// AuthorList
func (c *Controller) AuthorList(w http.ResponseWriter, req *http.Request) {

	results, err := service.GetAuthorsWithBooks(c.DB)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}

// AuthorGet
func (c *Controller) AuthorGet(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	fmt.Println(id)

	results, err := service.GetByIDWithBooks(c.DB, id)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}

// AuthorDelete
func (c *Controller) AuthorDelete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	err := service.DeleteAuthorByID(c.DB, id)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)

}
