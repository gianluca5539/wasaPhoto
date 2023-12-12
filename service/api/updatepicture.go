package api

import (
	"encoding/base64"
	"image/png"
	"bytes"
	"os"
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/gianluca5539/WASA/service/types"
)

type PictureRequest struct {
	NewPicture string `json:"newpicture"`
}


func (rt *_router) updatePicture(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	requestedUserID, err := strconv.Atoi(ps.ByName("id"))

	// get the new picture from the request body
	var pictureReq PictureRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&pictureReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}
	// get newPicture from pictureReq
	newPicture := pictureReq.NewPicture	

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
		errorobj := types.Error{Message: "You cannot update another user's picture"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// Decode the base64 string to []byte
	imgBytes, err := base64.StdEncoding.DecodeString(newPicture)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid picture"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	// Make a Reader out of the []byte
	imgReader := ioutil.NopCloser(bytes.NewReader(imgBytes))

	// Decode the []byte to image.Image
	img, err := png.Decode(imgReader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid picture"}
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

	// update the user's picture in the database
	err = rt.db.UpdatePicture(userID, timeInt)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		errorobj := types.Error{Message: "Internal server error"}
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	
	// return 201 created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// return the pictureid in the response body
	w.Write([]byte(`{"pictureid": "` + time + `"}`))
}