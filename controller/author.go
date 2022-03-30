package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/service"
	"github.com/gorilla/mux"
)

// controller
func (c *Controller) AuthorList(w http.ResponseWriter, req *http.Request) {

	results, err := service.GetAuthorsWithBooks(c.DB)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}

// controller
func (c *Controller) AuthorGet(w http.ResponseWriter, req *http.Request) {
	fmt.Println("aaaa")
	vars := mux.Vars(req)
	id := vars["id"]

	id_, _ := strconv.Atoi(id)

	fmt.Println(id)

	results, err := service.GetByIDWithBooks(c.DB, id_)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, results)

}
