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

// ListBooks controller for list all books
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

	limit, _ := strconv.ParseInt(args.Limit, 10, 64)
	offset, _ := strconv.ParseInt(args.Offset, 10, 64)

	data.Limit = limit
	data.Offset = offset

	RespondWithJSON(w, http.StatusOK, data)

}

// GetBook controller get book by id
func (c *Controller) GetBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	result, err := service.GetBookByIDWithAuthor(c.DB, uint(id))

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, result)

}

// DeleteBook controller delete book by id
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

// CreateBook controller create new book
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

// UpdateBook controller update book by id
func (c *Controller) UpdateBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	book := new(model.Book)

	if err := json.NewDecoder(req.Body).Decode(&book); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err = service.GetBookByID(c.DB, id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	defer req.Body.Close()

	book.ID = uint(id)

	book, err = service.UpdateBookByID(c.DB, book)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, book)

}

// BuyBook controller buy book by id with quantity payload
func (c *Controller) BuyBook(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}

	quantity := new(model.Quantity)

	if err := json.NewDecoder(req.Body).Decode(&quantity); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer req.Body.Close()

	instance, err := service.GetBookByID(c.DB, id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	if instance.StockCount < quantity.Amount {
		RespondWithError(w, http.StatusBadRequest, "Not enough quantity")
		return
	}
	fmt.Println(instance.StockCount)
	fmt.Println(quantity.Amount)

	remain := instance.StockCount - quantity.Amount

	if err := service.UpdateBookStockCount(c.DB, &instance, remain); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, instance)

}
