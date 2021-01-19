package handlers

import (
	"encoding/json"
	"github.com/fredalbert37/golang-rest-api/database"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func UpdateTodo( db database.TodoInterface) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var todo interface{}
		err = json.Unmarshal(body, &todo)
		if err != nil{
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := db.Update(id, todo)
		if err != nil{
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w,http.StatusOK, res)
	}
}