package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"golang_startup/internal/model"
	"golang_startup/internal/model/request"

	"github.com/gorilla/mux"
)

const (
	jsonContentType = "application/json"

	userDeletedMessage = "user successfully removed"
)

// GetUser - Handles a get user request.
// Requires a URL parameter ID.
// Returns a user ID or an error in case of failure.
func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := c.service.FindUser(context.Background(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// NewUser - Handles a user creation request.
// Requires a user data JSON body.
// Returns the created user or an error in case of failure.
func (c *Controller) NewUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	returnUser, err := c.service.CreateUser(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(returnUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// UpdateUser - Handles a user update request.
// Requires the user ID to update and a user data JSON body.
// Returns the updated user or an error in case of failure.
func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var user request.UserUpdate

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	updatedUser, err := c.service.UpdateUser(context.Background(), userID, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// DeleteUser - Handles a user deletion.
// Requires the user ID to be removed.
// Returns OK or an error in case of failure.
func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := c.service.DeleteUser(context.Background(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(userDeletedMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// FindUsers - Gets all users using pagination
// Optionally should be sent offset or limit pagination values.
// Returns OK or an error in case of failure.
func (c *Controller) FindUsers(w http.ResponseWriter, r *http.Request) {
	offset, limit, err := getPaginationValues(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := c.service.FindAllUsers(context.Background(), offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getPaginationValues(r *http.Request) (int64, int64, error) {
	const (
		emptyValue int64 = 0

		base    = 10
		bitSize = 64
	)

	rawOffset := r.URL.Query().Get("offset")
	rawLimit := r.URL.Query().Get("limit")

	offset, err := strconv.ParseInt(rawOffset, base, bitSize)
	if err != nil {
		return emptyValue, emptyValue, err
	}

	limit, err := strconv.ParseInt(rawLimit, base, bitSize)
	if err != nil {
		return emptyValue, emptyValue, err
	}

	return offset, limit, nil
}
