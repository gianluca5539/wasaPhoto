package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)



func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err := fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
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

	// get the user's following
	followingIds, err := rt.db.GetFollowing(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}


	// get stream from db
	var stream []types.Post
	stream, err = rt.db.GetStream(followingIds)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// make a response object
	res := struct {
		Posts []types.Post `json:"posts"`
	}{Posts: stream}
	
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)

}