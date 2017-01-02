package main

type ActionRequest struct {
	Name     string `json:"name,omitempty"`
	NewContext map[string]string `json:"newcontext,omitempty"`
	Context  map[string]string `json:"context,omitempty"`
}

type ActionResponse struct {
	Message string `json:"message,omitempty"`
	AddContext map[string]string `json:"addcontext,omitempty"`
	RemoveContext []string `json:"removecontext,omitempty"`
	E       error `json:"error,omitempty"`
	ReplaceContext map[string]string `json:"replacecontext,omitempty"`
}