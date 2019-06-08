package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go_sample/entity"
	user "go_sample/service"

	"github.com/go-chi/chi"
)

// Controller is user controlller
type Controller struct{}

// User is alias of entity.User struct
type User entity.User

// GET /users
func (c Controller) Index(w http.ResponseWriter, r *http.Request) {
	var u user.Service
	response, err := u.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GET /users/:userID
func (c Controller) Show(w http.ResponseWriter, r *http.Request) {
	var u user.Service

	userID := chi.URLParam(r, "userID")
	response, err := u.GetByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Create action: POST /users
func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	var u user.Service
	var user User

	// TODO: validation
	json.NewDecoder(r.Body).Decode(&user)

	response, err := u.CreateUser(user.FirstName, user.LastName)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Update action: PUT /users/:id
func (c Controller) Update(w http.ResponseWriter, r *http.Request) {
	var u user.Service
	var user User

	userID := chi.URLParam(r, "userID")
	// TODO: validation
	json.NewDecoder(r.Body).Decode(&user)

	response, err := u.UpdateByID(userID, user.FirstName, user.LastName)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DELETE /users/:id
func (c Controller) Delete(w http.ResponseWriter, r *http.Request) {
	var u user.Service

	userID := chi.URLParam(r, "userID")
	err := u.DeleteByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
