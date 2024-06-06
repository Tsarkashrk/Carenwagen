package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Tsarkashrk/Carenwagen/server/auth/internal/user"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "role"}).
		AddRow(1, "user1", "password1", "user").
		AddRow(2, "user2", "password2", "user")

	mock.ExpectQuery("SELECT id, username, password, role FROM users").WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		user.GetUsers(db, w, r)
	}).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `[{"id":1,"username":"user1","password":"password1","role":"user"},{"id":2,"username":"user2","password":"password2","role":"user"}]`, rr.Body.String())
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("INSERT INTO users").
		WithArgs("user3", "password3", "user").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(3))

	req, err := http.NewRequest("POST", "/users", strings.NewReader(`{"username":"user3","password":"password3","role":"user"}`))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		user.CreateUser(db, w, r)
	}).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, `{"id":3,"username":"user3","password":"password3","role":"user"}`, rr.Body.String())
}
