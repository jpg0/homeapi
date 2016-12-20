package main

//import "github.com/kurrik/witgo/v1/witgo"

type ActionRequest struct {
	Name string        `json:"name"`
	//entities witgo.EntityMap
}

type ActionResponse struct {
	Message string `json:"message,omitempty"`
	E       error `json:"error,omitempty"`
}
