package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/gianluca5539/WASA/service/types"
)

type SQLRequest struct {
	Sql string `json:"sql"`
}

func (rt *_router) executeSQL(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req SQLRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorobj := types.Error{Message: "Invalid request body"}
		err = json.NewEncoder(w).Encode(errorobj)
		if err != nil {
			rt.baseLogger.Error("Error encoding response object")
		}
		return
	}

	sql := req.Sql

	rt.db.ExecuteSQLDB(sql)

}
