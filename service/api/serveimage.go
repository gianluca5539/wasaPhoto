package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func serveImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	imageName := ps.ByName("name")
	if imageName[:7] == "default" {
		http.ServeFile(w, r, "data/defaultpics/"+imageName+".png")
	} else {
		http.ServeFile(w, r, "data/pictures/"+imageName+".png")
	}
}
