package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/gianluca5539/WASA/service/types"
)

type BioRequest struct {
	NewBio string `json:"newbio"`
}


func (rt *_router) updateBio(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID, err := strconv.Atoi(ps.ByName("id"))

	// get the new bio from the request body
	var bioReq BioRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&bioReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	// get newBio from bioReq
	newBio := bioReq.NewBio	

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

	if requestedUserID != userID {
		w.WriteHeader(http.StatusForbidden)
		errorobj := types.Error{Message: "You cannot update another user's bio"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}


	// update the user in the database
	dberr := rt.db.UpdateBio(userID, newBio)
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