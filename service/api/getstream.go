package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err := fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "Invalid token"}
		rt.baseLogger.Error("Invalid user token")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// convert the token string to an int
	userID, err := strconv.Atoi(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "Invalid token"}
		rt.baseLogger.Error("Invalid user token")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// get the user's following
	followingIds, err := rt.db.GetFollowing(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		rt.baseLogger.Error("Error getting following from db", err)
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object", err)
		}
		return
	}

	// get stream from db
	var stream []types.Post
	stream, err = rt.db.GetStream(followingIds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		rt.baseLogger.Error("Error getting stream from db", err)
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object", err)
		}
		return
	}

	// make a response object
	res := struct {
		Posts []types.Post `json:"posts"`
	}{Posts: stream}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		rt.baseLogger.Error("Error encoding response object", err)
	}

}
