package router

import (
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitializeRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	handlers.AllowedOrigins([]string{"*"})
	handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	api := controller.Controller{DB: db}
	// routers
	author := r.PathPrefix("/authors").Subrouter()
	author.HandleFunc("/", api.AuthorList).Methods(http.MethodGet)
	// author.HandleFunc("/", authorCreate).Methods(http.MethodPost)
	author.HandleFunc("/{id}", api.AuthorGet).Methods(http.MethodGet)
	// author.HandleFunc("/{id}", authorUpdate).Methods(http.MethodPut)
	// author.HandleFunc("/{id}", authorDelete).Methods(http.MethodDelete)

	book := r.PathPrefix("/books").Subrouter()
	// book.HandleFunc("/", bookCreate).Methods(http.MethodPost)
	book.HandleFunc("/", api.BookList).Methods(http.MethodGet)
	book.HandleFunc("/{id}", api.BookGet).Methods(http.MethodGet)
	// book.HandleFunc("/{id}", bookUpdate).Methods(http.MethodPut)
	// book.HandleFunc("/{id}", bookDelete).Methods(http.MethodDelete)
	return r
}
