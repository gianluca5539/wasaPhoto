package api

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)



func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID , err := strconv.Atoi(ps.ByName("id"))

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

	// get the user profile from the database
	banned, err := rt.db.IsUserBanned(userID, requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if banned {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check the user is not following themselves
	if userID == requestedUserID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.FollowUser(requestedUserID, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	// return the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}