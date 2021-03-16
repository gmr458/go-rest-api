package main

import (
	"fmt"
	"net/http"

	"github.com/gmr458/go-rest-api/src/api"
	"github.com/gorilla/mux"
)

func main() {
	var port string = "8080"
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/home", api.Home).Methods("GET")
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")
	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
