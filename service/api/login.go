package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

type LoginRequest struct {
	Username string `json:"username"`
}


func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var loginReq LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := loginReq.Username

	// Check if the username exists
	user,found, err := rt.db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If the user doesn't exist, create a new user in the database
	if !found {
		user, err = rt.db.CreateUser(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} 
		token := user.UserID

		res := map[string]interface{}{
			"user": user,
			"token": token,
		}
		// Return the user and the token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	} else {
		token := user.UserID

		res := map[string]interface{}{
			"user": user,
			"token": token,
		}
		// Return the user and the token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}; 
