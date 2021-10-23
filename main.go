package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHome)
	log.Println("Application running at http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello from Golang"}`))

}
