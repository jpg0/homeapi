package main

import (
	"net/http"
)


func setupRouting(port string, cfg map[string]string) {
	http.HandleFunc("/", ConfigureHandleAction(cfg))
	err := http.ListenAndServe(":" + port, nil)

	if err != nil {
		panic(err)
	}
}