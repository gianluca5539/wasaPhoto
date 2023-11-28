package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// getHelloWorld is an example of HTTP endpoint that returns "Hello world!" as a plain text
func (rt *_router) getDeleteData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rt.db.DeleteAllData()
	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte("done"))
}
