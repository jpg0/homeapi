package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	setupHandlers();
	setupRouting();
}

func setupRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/runAction", HandleAction)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func setupHandlers() {
	InitListActions();
	InitRestartActions();
}