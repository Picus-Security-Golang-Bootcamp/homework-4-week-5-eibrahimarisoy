package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/model"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/service"
	"github.com/gorilla/mux"
)

func (c *Controller) ListBook(w http.ResponseWriter, req *http.Request) {
	args := model.Args{Sort: "ID", Order: "desc", Offset: "0", Limit: "10", Search: ""}

	v := req.URL.Query()

	if v["Sort"] != nil {
		args.Sort = v.Get("Sort")
	}

	if v["Limit"] != nil {
		args.Limit = v.Get("Limit")
	}

	if v["Offset"] != nil {
		args.Offset = v.Get("Offset")
	}

	if v["Search"] != nil {
		args.Search = v.Get("Search")
	}

	data, err := service.GetBooksWithAuthor(c.DB, args)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, data)

}

// controller
func (c *Controller) GetBook(w http.ResponseWriter, req *http.Request) {
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
func (c *Controller) DeleteBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	if err := service.DeleteBookByID(c.DB, id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusNoContent, nil)

}

// CreatePost controller
func (c *Controller) CreateBook(w http.ResponseWriter, req *http.Request) {
	book := new(model.Book) // model.Book{}

	err := json.NewDecoder(req.Body).Decode(&book)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer req.Body.Close()

	book, err = service.SaveBook(c.DB, book)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, book)

}

// UpdatePost controller
func (c *Controller) UpdateBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	newBook := new(model.Book) // model.Book{}

	if err := json.NewDecoder(req.Body).Decode(&newBook); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err = service.GetBookByID(c.DB, id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	defer req.Body.Close()

	newBook.ID = uint(id)

	newBook, err = service.UpdateBookByID(c.DB, newBook)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, newBook)

}
