package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	handlers.AllowedOrigins([]string{"*"})
	handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	srv := &http.Server{
		Addr:         ":8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	// routers
	a := r.PathPrefix("/authors").Subrouter()
	a.HandleFunc("/", authorList).Methods(http.MethodGet)
	// a.HandleFunc("/", authorCreate).Methods(http.MethodPost)
	// a.HandleFunc("/{id}", authorGet).Methods(http.MethodGet)
	// a.HandleFunc("/{id}", authorUpdate).Methods(http.MethodPut)
	// a.HandleFunc("/{id}", authorDelete).Methods(http.MethodDelete)

	// // routes
	// b := r.PathPrefix("/books").Subrouter()
	// b.HandleFunc("/", bookCreate).Methods(http.MethodPost)
	// b.HandleFunc("/", bookList).Methods(http.MethodGet)
	// b.HandleFunc("/{id}", bookGet).Methods(http.MethodGet)
	// b.HandleFunc("/{id}", bookUpdate).Methods(http.MethodPut)
	// b.HandleFunc("/{id}", bookDelete).Methods(http.MethodDelete)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ShutdownServer(srv, time.Second*10)

}

func authorList(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")

}

//https://medium.com/@pinkudebnath/graceful-shutdown-of-golang-servers-using-context-and-os-signals-cc1fa2c55e97
//https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/
func ShutdownServer(srv *http.Server, timeout time.Duration) {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
