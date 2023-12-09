package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

type UsernameRequest struct {
	NewUsername string `json:"newusername"`
}


func (rt *_router) updateUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID, err := strconv.Atoi(ps.ByName("id"))

	// get the new username from the request body
	var usernameReq UsernameRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&usernameReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get newUsername from usernameReq
	newUsername := usernameReq.NewUsername	

	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err = fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert the token string to an int
	userID, err := strconv.Atoi(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if there is another user with the same username
	_,found, err := rt.db.GetUserByUsername(newUsername)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if found {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if requestedUserID != userID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// check new username is valid (3 to 16 characters)
	if len(newUsername) < 3 || len(newUsername) > 16 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// update the user in the database
	dberr := rt.db.UpdateUsername(userID, newUsername)
	if dberr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	
	// return 204 no content
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}