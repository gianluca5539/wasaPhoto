package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
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
		_ = json.NewEncoder(w).Encode(errorobj)
		return
	}

	sql := req.Sql
	fmt.Println("executing sql code: " + sql)

	rt.db.ExecuteSQLDB(sql)
	
}; 
