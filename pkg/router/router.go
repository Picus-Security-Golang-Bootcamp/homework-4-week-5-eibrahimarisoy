package router

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// InitializeRoutes initializes the routes for the application
func InitializeRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	handlers.AllowedOrigins([]string{"*"})
	handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	api := controller.Controller{DB: db}
	// routers
	author := r.PathPrefix("/authors").Subrouter()
	author.HandleFunc("", api.ListAuthor).Methods(http.MethodGet)
	author.HandleFunc("", api.CreateAuthor).Methods(http.MethodPost)
	author.HandleFunc("/{id}", api.GetAuthor).Methods(http.MethodGet)
	author.HandleFunc("/{id}", api.UpdateAuthor).Methods(http.MethodPut)
	author.HandleFunc("/{id}", api.DeleteAuthor).Methods(http.MethodDelete)
	author.HandleFunc("/{id}/books", api.ListAuthorsBooks).Methods(http.MethodGet)

	book := r.PathPrefix("/books").Subrouter()
	book.HandleFunc("", api.CreateBook).Methods(http.MethodPost)
	book.HandleFunc("", api.ListBook).Methods(http.MethodGet)
	book.HandleFunc("/{id}", api.GetBook).Methods(http.MethodGet)
	book.HandleFunc("/{id}", api.UpdateBook).Methods(http.MethodPut)
	book.HandleFunc("/{id}", api.DeleteBook).Methods(http.MethodDelete)
	book.HandleFunc("/{id}/buy", api.BuyBook).Methods(http.MethodPost)

	return r
}
