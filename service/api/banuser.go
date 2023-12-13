package api

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/gianluca5539/WASA/service/types"
)



func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID , err := strconv.Atoi(ps.ByName("id"))

	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err = fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid token"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// convert the token string to an int
	userID, err := strconv.Atoi(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid token"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// check the user is not banning themselves
	if userID == requestedUserID {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "You cannot ban yourself"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	err = rt.db.BanUser(requestedUserID, userID)
	if err != nil {
		if err.Error() == "User is already banned" {
			w.WriteHeader(http.StatusBadRequest)
			errorobj := types.Error{Message: "User is already banned"}
			_ = json.NewEncoder(w).Encode(errorobj)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	
	// return the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}