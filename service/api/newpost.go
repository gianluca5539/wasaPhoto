package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

type postRequest struct {
	Image   string `json:"image"`
	Caption string `json:"caption"`
}

func (rt *_router) newPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var postReq postRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
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

	image := postReq.Image
	caption := postReq.Caption

	// save image locally
	// Decode the base64 string to []byte
	imgBytes, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid image"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// Make a Reader out of the []byte
	imgReader := ioutil.NopCloser(bytes.NewReader(imgBytes))

	// Decode the []byte to image.Image
	img, err := png.Decode(imgReader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid image"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// Create the output file with the current timestamp as the filename
	currentTime := time.Now()
	time := strconv.FormatInt(currentTime.UnixNano()/int64(time.Millisecond), 10)
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	// save picture
	filename := "data/pictures/" + time + ".png"
	outputFile, err := os.Create(filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	defer outputFile.Close()

	// Encode the image to the output file
	err = png.Encode(outputFile, img)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// insert the post into the database (userid,imageid,caption,timestamp)
	postid, err := rt.db.CreateNewPost(userID, timeInt, caption, timeInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// if everything is successful, get the user object to return to the client together with the post
	u, _, err := rt.db.GetUserByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	res := types.Post{
		PostID:      postid,
		UserID:      u.UserID,
		Username:    u.Username,
		Feeling:     u.Feeling,
		UserPicture: u.Picture,
		Picture:     timeInt,
		Caption:     caption,
		CreatedAt:   timeInt,
		LikeCount:   0,
	}

	w.WriteHeader(http.StatusCreated)
	// return the user object to the client
	json.NewEncoder(w).Encode(res)

}
