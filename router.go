package main

import (
	"github.com/gorilla/mux"
	"net/http"
)


func setupRouting(port string, cfg map[string]string) {
	r := mux.NewRouter()
	r.HandleFunc("/runAction", ConfigureHandleAction(cfg))
	http.Handle("/", r)
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		panic(err)
	}
}