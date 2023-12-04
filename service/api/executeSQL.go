package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
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
		return
	}

	sql := req.Sql
	fmt.Println("executing sql code: " + sql)

	rt.db.ExecuteSQLDB(sql)
	
}; 
