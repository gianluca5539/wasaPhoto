package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

type CommentRequest struct {
	Text string `json:"text"`
}

func (rt *_router) commentPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	postID, err := strconv.Atoi(ps.ByName("postid"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid post id"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// get the text from the request body
	var commentRequest CommentRequest
	err = json.NewDecoder(r.Body).Decode(&commentRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	text := commentRequest.Text

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

	currentTime := time.Now()
	time := strconv.FormatInt(currentTime.UnixNano()/int64(time.Millisecond), 10)
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	commentID, err := rt.db.CreateComment(postID, userID, text, timeInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// add user and create response
	u, _, err := rt.db.GetUserByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	comment := types.Comment{
		ID:        commentID,
		UserID:    u.UserID,
		PostID:    postID,
		Username:  u.Username,
		Picture:   u.Picture,
		Feeling:   u.Feeling,
		Text:      text,
		CreatedAt: timeInt,
	}

	// return the user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}
