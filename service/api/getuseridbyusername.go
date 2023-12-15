package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

func (rt *_router) getUserIDByUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := ps.ByName("username")

	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err := fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "Invalid token"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	userID, err := strconv.Atoi(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "Invalid token"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// get user by username
	user, found, err := rt.db.GetUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// check if the user is banned
	banned, err := rt.db.IsUserBanned(userID, user.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	if !found || banned {
		w.WriteHeader(http.StatusNotFound)
		errorobj := types.Error{Message: "User not found"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// make a response object
	res := struct {
		ID int `json:"id"`
	}{ID: user.UserID}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)

}