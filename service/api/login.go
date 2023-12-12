package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gianluca5539/WASA/service/types"
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
		errorobj := types.Error{Message: "Invalid request body"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	username := loginReq.Username

	// check username is valid (3 to 16 characters)
	if len(username) < 3 || len(username) > 16 {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid username"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// Check if the username exists
	user,found, err := rt.db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	// If the user doesn't exist, create a new user in the database
	if !found {
		user, err = rt.db.CreateUser(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorobj := types.Error{Message: "Internal server error"}
			_ = json.NewEncoder(w).Encode(errorobj)
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
