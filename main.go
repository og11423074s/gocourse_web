package main

import (
	"github.com/joho/godotenv"
	"github.com/og11423074s/go_course_web/pkg/bootstrap"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/og11423074s/go_course_web/internal/user"
)

func main() {

	router := mux.NewRouter()
	// Load .env file
	_ = godotenv.Load()

	// Initialize logger
	logger := bootstrap.InitLogger()

	// Connect to database
	db, err := bootstrap.DBConnection()

	if err != nil {
		logger.Fatal(err)
	}

	// User repository
	userRepo := user.NewRepo(logger, db)
	// User service
	userSrv := user.NewService(logger, userRepo)
	userEnd := user.MakeEndpoints(userSrv)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userEnd.Get).Methods("GET")
	router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      http.TimeoutHandler(router, time.Second*3, "Timeout!!"),
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
