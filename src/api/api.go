package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gmr458/go-rest-api/src/helpers"
	"github.com/gmr458/go-rest-api/src/models"
	"github.com/gorilla/mux"
)

type Data struct {
	Success bool          `json:"success"`
	Data    []models.Todo `json:"data"`
	Errors  []string      `json:"errors"`
}

type Message struct {
	Message string `json:"message"`
}

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	bodyTodo, success := helpers.DecodeBody(req)
	if !success {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	todo, success := models.Insert(bodyTodo.Description)
	if !success {
		data.Errors = append(data.Errors, "could not create todo")
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

func GetTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var data Data
	var todo models.Todo
	var success bool
	todo, success = models.Get(id)
	if !success {
		data.Success = false
		data.Errors = append(data.Errors, "todo not found")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func Home(w http.ResponseWriter, req *http.Request) {
	var res Message
	res.Message = "Home"
	json, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
