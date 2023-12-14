package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gianluca5539/WASA/service/types"
	"github.com/julienschmidt/httprouter"
)

type FeelingRequest struct {
	NewFeeling int `json:"newfeeling"`
}


func (rt *_router) updateFeeling(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid user id"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// get the new feeling from the request body
	var feelingReq FeelingRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&feelingReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	// get newFeeling from feelingReq
	newFeeling := feelingReq.NewFeeling	

	// check that newFeeling is valid (between 0 and 4)
	if newFeeling < 0 || newFeeling > 4 {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid feeling"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

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
		errorobj := types.Error{Message: "You cannot update another user's feeling"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// update the user in the database
	dberr := rt.db.UpdateFeeling(userID, newFeeling)
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