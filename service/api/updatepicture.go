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
		return
	}
	// get newPicture from pictureReq
	newPicture := pictureReq.NewPicture	

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

	// Decode the base64 string to []byte
	imgBytes, err := base64.StdEncoding.DecodeString(newPicture)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Make a Reader out of the []byte
	imgReader := ioutil.NopCloser(bytes.NewReader(imgBytes))

	// Decode the []byte to image.Image
	img, err := png.Decode(imgReader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create the output file with the current timestamp as the filename
	currentTime := time.Now()
	time := strconv.FormatInt(currentTime.UnixNano()/int64(time.Millisecond), 10)
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// save picture
	filename := "data/pictures/" + time + ".png"
	outputFile, err := os.Create(filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Encode the image to the output file
	err = png.Encode(outputFile, img)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// update the user's picture in the database
	err = rt.db.UpdatePicture(userID, timeInt)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	
	// return 204 no content
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}