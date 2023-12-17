package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

func (rt *_router) deletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	postID, err := strconv.Atoi(ps.ByName("postid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid postid"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

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

	// if everything is successful, get the user object to return to the client together with the post
	err = rt.db.DeletePostCascading(userID, postID)
	if err != nil {
		// check if err = "you are not the owner of the post"
		if err.Error() == "you are not the owner of the post" {
			w.WriteHeader(http.StatusForbidden)
			errorobj := types.Error{Message: "you are not the owner of the post"}
			_ = json.NewEncoder(w).Encode(errorobj)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
