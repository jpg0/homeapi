package main

import "github.com/jpg0/witgo/v1/witgo"

type ActionRequest struct {
	Name     string `json:"name,omitempty"`
	Entities witgo.EntityMap `json:"entities,omitempty"`
	Context  map[string]string `json:"context,omitempty"`
}

type ActionResponse struct {
	Message string `json:"message,omitempty"`
	Context map[string]string `json:"context,omitempty"`
	E       error `json:"error,omitempty"`
}