package handlers

import (
	"github.com/fredalbert37/golang-rest-api/database"
	"net/http"
)

func ListAllTodos(db database.TodoInterface) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){

	}

}