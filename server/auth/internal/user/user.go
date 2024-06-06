package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tsarkashrk/Carenwagen/server/auth/pkg/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetUsers(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var users []models.User
	rows, err := db.Query("SELECT id, username, password, role FROM users")
	if err != nil {
		logrus.Error("Failed to execute query: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role); err != nil {
			logrus.Error("Failed to scan row: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logrus.Error("Failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, user.Username, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		logrus.Error("Failed to execute query: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Error("Invalid user ID: ", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logrus.Error("Failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET username=$1, password=$2, role=$3 WHERE id=$4`
	_, err = db.Exec(query, user.Username, user.Password, user.Role, userID)
	if err != nil {
		logrus.Error("Failed to execute query: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Error("Invalid user ID: ", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM users WHERE id=$1`
	_, err = db.Exec(query, userID)
	if err != nil {
		logrus.Error("Failed to execute query: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
