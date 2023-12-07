package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

type FeelingRequest struct {
	NewFeeling int `json:"newfeeling"`
}


func (rt *_router) updateFeeling(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID, err := strconv.Atoi(ps.ByName("id"))

	// get the new feeling from the request body
	var feelingReq FeelingRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&feelingReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get newFeeling from feelingReq
	newFeeling := feelingReq.NewFeeling	

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

	if requestedUserID != userID {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// update the user in the database
	dberr := rt.db.UpdateFeeling(userID, newFeeling)
	if dberr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	// return 204 no content
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}