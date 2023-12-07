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

	// convert the token string to an int
	userID, err := strconv.Atoi(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	// find the followers ids of the user that is requested
	followersIds, err := rt.db.GetFollowers(requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the user objects of the followers
	var followers []User
	for _, id := range followersIds {
		u, _, err := rt.db.GetUserByUserID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		followers = append(followers, User{UserID: u.UserID, Username: u.Username, Picture: u.Picture, Bio: u.Bio, Feeling: u.Feeling})
	}

	// find the following ids of the user that is requested
	followingIds, err := rt.db.GetFollowing(requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the user objects of the following
	var following []User
	for _, id := range followingIds {
		u, _, err := rt.db.GetUserByUserID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		following = append(following, User{UserID: u.UserID, Username: u.Username, Picture: u.Picture, Bio: u.Bio, Feeling: u.Feeling})
	}



	

	// create the response
	res := struct {
		ID int `json:"id"`
		Username string `json:"username"`
		Picture string `json:"picture"`
		Feeling int `json:"feeling"`
		Bio string `json:"bio"`
		Followers []User `json:"followers"`
		Following []User `json:"following"`
	}{
		ID: u.UserID,
		Username: u.Username,
		Picture: u.Picture,
		Feeling: u.Feeling,
		Bio: u.Bio,
		Followers: followers,
		Following: following,
	}

	// return the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}