package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

type UsernameRequest struct {
	NewUsername string `json:"newusername"`
}

func (rt *_router) updateUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid user id"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// get the new username from the request body
	var usernameReq UsernameRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&usernameReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	// get newUsername from usernameReq
	newUsername := usernameReq.NewUsername

	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err = fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "Invalid token"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// convert the token string to an int
	userID, err := strconv.Atoi(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "Invalid token"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// check if there is another user with the same username
	_, found, err := rt.db.GetUserByUsername(newUsername)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	if found {
		w.WriteHeader(http.StatusConflict)
		errorobj := types.Error{Message: "Username already taken"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	if requestedUserID != userID {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "You cannot update another user's username"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// check new username is valid (3 to 16 characters)
	if len(newUsername) < 3 || len(newUsername) > 16 {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid username"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// update the user in the database
	dberr := rt.db.UpdateUsername(userID, newUsername)
	if dberr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// return 204 no content
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
