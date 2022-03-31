package controller

import (
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/service"
	"github.com/gorilla/mux"
)

func (c *Controller) BookList(w http.ResponseWriter, req *http.Request) {

	results, err := service.GetBooksWithAuthor(c.DB)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}

// controller
func (c *Controller) BookGet(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	results, err := service.GetByIDWithAuthor(c.DB, id)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}

// BookDelete controller
func (c *Controller) BookDelete(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.Atoi(vars["id"])

	err := service.DeleteBookByID(c.DB, id)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)

}
