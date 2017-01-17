package main

import (
	"github.com/Sirupsen/logrus"
)

type ActionHandler interface {
	Run(req *APIAIRequest, config *Configuration) (*APIAIResponse, error)
}

var handlerFactory = make(map[string]ActionHandler)

func RegisterHandler(name string, handler ActionHandler) {
	if handler == nil {
		logrus.Fatalf("Datastore factory %s does not exist.", name)
	}
	_, registered := handlerFactory[name]
	if registered {
		logrus.Fatalf("Datastore factory %s already registered. Ignoring.", name)
	}

	handlerFactory[name] = handler
}

func GetHandler(name string) ActionHandler {
	logrus.Infof("Loading handler: %v", name)
	return handlerFactory[name]
}