package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

func (rt *_router) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid user id"}
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

	// get the user profile from the database
	banned, err := rt.db.IsUserBanned(userID, requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}
	if banned {
		w.WriteHeader(http.StatusNotFound)
		errorobj := types.Error{Message: "User not found"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// do it in reverse to check if the user is banned by the requested user
	bannedByReqUser, err := rt.db.IsUserBanned(requestedUserID, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// find the user that is requested
	// get the user profile from the database
	u, found, err := rt.db.GetUserByUserID(requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		errorobj := types.Error{Message: "User not found"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// find the followers ids of the user that is requested
	followersIds, err := rt.db.GetFollowers(requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// get the user objects of the followers
	var followers []types.User
	for _, id := range followersIds {
		u, _, err := rt.db.GetUserByUserID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorobj := types.Error{Message: "Internal server error"}
			err = json.NewEncoder(w).Encode(errorobj)
			if err != nil {
				rt.baseLogger.Error("Error encoding response object")
			}
			return
		}
		followers = append(followers, types.User{UserID: u.UserID, Username: u.Username, Picture: u.Picture, Bio: u.Bio, Feeling: u.Feeling})
	}

	// find the following ids of the user that is requested
	followingIds, err := rt.db.GetFollowing(requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// get the user objects of the following
	var following []types.User
	for _, id := range followingIds {
		u, _, err := rt.db.GetUserByUserID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errorobj := types.Error{Message: "Internal server error"}
			err = json.NewEncoder(w).Encode(errorobj)
			if err != nil {
				rt.baseLogger.Error("Error encoding response object")
			}
			return
		}
		following = append(following, types.User{UserID: u.UserID, Username: u.Username, Picture: u.Picture, Bio: u.Bio, Feeling: u.Feeling})
	}

	// get the user's posts
	posts, err := rt.db.GetPostsByUserID(requestedUserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// create the response
	res := types.UserProfile{
		ID:        u.UserID,
		Username:  u.Username,
		Picture:   u.Picture,
		Feeling:   u.Feeling,
		Bio:       u.Bio,
		Followers: followers,
		Following: following,
		Banned:    bannedByReqUser,
		Posts:     posts,
	}

	// return the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		rt.baseLogger.Error("Error encoding response object")
	}
}
