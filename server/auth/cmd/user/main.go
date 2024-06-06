package main

import (
	"log"
	"net/http"

	"github.com/Tsarkashrk/Carenwagen/server/auth/config"
	"github.com/Tsarkashrk/Carenwagen/server/auth/internal/database"
	"github.com/Tsarkashrk/Carenwagen/server/auth/internal/user"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file: ", err)
	}

	// Load config from environment variables
	cfg := config.LoadConfig()
	logrus.Infof("Loaded config: %+v", cfg)

	// Initialize database
	db, err := database.InitDB(&cfg)
	if err != nil {
		logrus.Fatal("Failed to initialize database: ", err)
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		user.GetUsers(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		user.CreateUser(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		user.UpdateUser(db, w, r)
	}).Methods("PUT")

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		user.DeleteUser(db, w, r)
	}).Methods("DELETE")

	logrus.Info("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
