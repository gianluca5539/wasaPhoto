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

func (rt *_router) newPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var postReq postRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		rt.baseLogger.Error("Invalid request body")
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
		rt.baseLogger.Error("Invalid user token")
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
		rt.baseLogger.Error("Invalid user token")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
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
		rt.baseLogger.Error("Invalid image")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// Make a Reader out of the []byte
	imgReader := ioutil.NopCloser(bytes.NewReader(imgBytes))

	// Decode the []byte to image.Image
	img, err := png.Decode(imgReader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid image"}
		rt.baseLogger.Error("Invalid image")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// Create the output file with the current timestamp as the filename
	currentTime := time.Now()
	time := strconv.FormatInt(currentTime.UnixNano()/int64(time.Millisecond), 10)
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		rt.baseLogger.Error("Error converting timestamp to int")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}
	// save picture
	filename := "data/pictures/" + time + ".png"
	outputFile, err := os.Create(filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		rt.baseLogger.Error("Error creating output file")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}
	defer outputFile.Close()

	// Encode the image to the output file
	err = png.Encode(outputFile, img)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		rt.baseLogger.Error("Error encoding image to output file")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// insert the post into the database (userid,imageid,caption,timestamp)
	postid, err := rt.db.CreateNewPost(userID, timeInt, caption, timeInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		rt.baseLogger.Error("Error inserting post into db")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	// if everything is successful, get the user object to return to the client together with the post
	u, _, err := rt.db.GetUserByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		rt.baseLogger.Error("Error getting user from db")
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
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
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		rt.baseLogger.Error("Error encoding response object")
	}

}
