package handlers

import (
	"github.com/fredalbert37/golang-rest-api/database"
	"github.com/gorilla/mux"
	"net/http"
)

func DeleteTodo( db database.TodoInterface) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		res, err := db.Delete(id)
		if err != nil{
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w,http.StatusOK, res)
	}
}