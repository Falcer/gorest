package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type todo struct {
	ID          int    `json:"id,omitempty"`          // 0
	Title       string `json:"title,omitempty"`       // ""
	Description string `json:"description,omitempty"` // ""
	IsDone      bool   `json:"is_done,omitempty"`     // false
}

// Mutable + Immutable
var todo_database = []todo{
	{
		ID:          1,
		Title:       "First Todo",
		Description: "this is my first todo",
		IsDone:      false,
	},
	{
		ID:          2,
		Title:       "Second Todo",
		Description: "this is my second todo",
		IsDone:      true,
	},
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/todo", getAllTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", getTodoByID).Methods("GET")
	r.HandleFunc("/todo", createTodo).Methods("POST")
	r.HandleFunc("/todo/{id}", updateTodoByID).Methods("PUT")
	r.HandleFunc("/todo", deleteTodoByID).Methods("DELETE")
	http.ListenAndServe(":8000", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(response{
		Message: "application is running",
	})
}

func getAllTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(response{
		Message: "Get all todos",
		Data:    todo_database,
	})
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"]) // success -> err == nil, failed -> err != nil
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Message: "{id} must be integer",
		})
	}
	var resTodo *todo // {}
	for _, item := range todo_database {
		if item.ID == id {
			resTodo = &item
			break
		}
	}
	json.NewEncoder(w).Encode(response{
		Message: fmt.Sprintf("Get todo by id %d", id),
		Data:    resTodo,
	})
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo
	err := json.NewDecoder(r.Body).Decode(&newTodo) // nil == success, not nil == failed
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Message: "Create todo failed",
		})
	}
	// Set ID auth
	newTodo.ID = todo_database[len(todo_database)-1].ID + 1
	// Set Default is_done to false, for simulate this todo is not done yet
	newTodo.IsDone = false
	// Add to (fake) database
	todo_database = append(todo_database, newTodo)
	// Send success response
	json.NewEncoder(w).Encode(response{
		Message: "Create todo successed",
		Data:    newTodo,
	})
}

func updateTodoByID(w http.ResponseWriter, r *http.Request) {
	// Slice Golang, update spesific element slice of struct golang
}

func deleteTodoByID(w http.ResponseWriter, r *http.Request) {
	// Slice Golang, remove spesific element slice of struct golang
}
