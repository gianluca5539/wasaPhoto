package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

func (rt *_router) unCommentPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	commentID, err := strconv.Atoi(ps.ByName("commentid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid post id"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err = fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "Invalid token"}
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
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	err = rt.db.RemoveComment(commentID, userID)
	if err != nil {
		if err.Error() == "comment does not exist" {
			w.WriteHeader(http.StatusNotFound)
			errorobj := types.Error{Message: "Comment does not exist or you are not the owner"}
			err = json.NewEncoder(w).Encode(errorobj)
			if err != nil {
				rt.baseLogger.Error("Error encoding response object")
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// return the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
