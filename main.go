package main

import (
	"context"
	"github.com/fredalbert37/golang-rest-api/config"
	"github.com/fredalbert37/golang-rest-api/database"
	"github.com/fredalbert37/golang-rest-api/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()
	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &database.TodoClient{
		Ctx: ctx,
		Col: collection,
	}

	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.SearchTodos(client)).Methods("GET")
	r.HandleFunc("/todos/{id}", handlers.GetTodo(client)).Methods("GET")
	r.HandleFunc("/todos", handlers.InsertTodo(client)).Methods("POST")
	r.HandleFunc("/todo/{id}", handlers.UpdateTodo(client)).Methods("PATCH")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo(client)).Methods("DELETE")

	_ = http.ListenAndServe(":8080", r)
}
