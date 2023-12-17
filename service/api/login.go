package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

// LoginRequest represents the request structure for the login API.
type LoginRequest struct {
	Username string `json:"username"`
}

func (rt *_router) login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var loginReq LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	username := loginReq.Username

	// check username is valid (3 to 16 characters)
	if len(username) < 3 || len(username) > 16 {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid username"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// Check if the username exists
	user, found, err := rt.db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}
	// If the user doesn't exist, create a new user in the database
	if !found {
		user, err = rt.db.CreateUser(username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorobj := types.Error{Message: "Internal server error"}
			err = json.NewEncoder(w).Encode(errorobj)
			if err != nil {
				rt.baseLogger.Error("Error encoding response object")
			}
			return
		}
		token := user.UserID

		res := map[string]interface{}{
			"user":  user,
			"token": token,
		}
		// Return the user and the token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
	} else {
		token := user.UserID

		res := map[string]interface{}{
			"user":  user,
			"token": token,
		}
		// Return the user and the token
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
	}
}
