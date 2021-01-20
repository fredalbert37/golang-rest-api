package handlers

import (
	"github.com/fredalbert37/golang-rest-api/database"
	"net/http"
)

func ListAllTodos(db database.TodoInterface) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){
		res, err := db.List()
		if err != nil{
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}

}