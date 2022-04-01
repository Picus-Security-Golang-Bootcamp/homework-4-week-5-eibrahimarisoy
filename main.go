package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/pkg/database"
	"github.com/Picus-Security-Golang-Bootcamp/homework-4-week-5-eibrahimarisoy/pkg/router"

	"github.com/joho/godotenv"
)

func main() {
	// Set environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	db, err := database.NewPsqlDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	r := router.InitializeRoutes(db)

	// Start server
	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	srv.ListenAndServe()

}
