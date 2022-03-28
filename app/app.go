package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/pkg/common/db"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/pkg/domain/repos"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router     *mux.Router
	authorRepo *repos.AuthorRepository
}

func (a *App) Initialize() error {
	db, err := db.NewPsqlDB()
	if err != nil {
		return nil
	}
	authorRepo := repos.NewAuthorRepository(db)
	a.authorRepo = authorRepo

	a.Router = mux.NewRouter()

	a.InitializeRoutes()
	return nil
}

func (a *App) Run(srv *http.Server) {
	srv.Handler = a.Router
	log.Fatal(srv.ListenAndServe())
}

func (a *App) InitializeRoutes() {
	handlers.AllowedOrigins([]string{"*"})
	handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// routers
	author := a.Router.PathPrefix("/authors").Subrouter()
	author.HandleFunc("/", a.authorList).Methods(http.MethodGet)
	// author.HandleFunc("/", authorCreate).Methods(http.MethodPost)
	// author.HandleFunc("/{id}", authorGet).Methods(http.MethodGet)
	// author.HandleFunc("/{id}", authorUpdate).Methods(http.MethodPut)
	// author.HandleFunc("/{id}", authorDelete).Methods(http.MethodDelete)

	// // routes
	// book := r.PathPrefix("/books").Subrouter()
	// book.HandleFunc("/", bookCreate).Methods(http.MethodPost)
	// book.HandleFunc("/", bookList).Methods(http.MethodGet)
	// book.HandleFunc("/{id}", bookGet).Methods(http.MethodGet)
	// book.HandleFunc("/{id}", bookUpdate).Methods(http.MethodPut)
	// book.HandleFunc("/{id}", bookDelete).Methods(http.MethodDelete)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) authorList(w http.ResponseWriter, req *http.Request) {

	results, err := a.authorRepo.GetAuthorsWithBooks()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, results)

}
