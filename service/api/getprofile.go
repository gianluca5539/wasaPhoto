package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)



func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID , err := strconv.Atoi(ps.ByName("id"))

	// get the user id from the jwt token in the request header (bearer token)
	var tokenString string
	_, err = fmt.Sscanf(r.Header.Get("Authorization"), "Bearer %s", &tokenString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// parse the jwt token
	token, err := jwt.ParseWithClaims(tokenString, &CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("wasaphoto_secret"), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if the token is valid
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// get the user id from the jwt claims
	claims, ok := token.Claims.(*CustomJWTClaims)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userID := claims.UserID

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

	// find the user that is requested
	// get the user profile from the database
	u, found, err := rt.db.GetUserByUserID(requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	

	// create the response
	res := struct {
		ID int `json:"id"`
		Username string `json:"username"`
		Picture string `json:"picture"`
		Feeling int `json:"feeling"`
	}{
		ID: u.UserID,
		Username: u.Username,
		Picture: u.Picture,
		Feeling: u.Feeling,
	}

	// return the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}