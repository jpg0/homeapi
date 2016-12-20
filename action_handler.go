package main

import (
	"net/http"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/juju/errors"
)

func HandleAction(w http.ResponseWriter, r *http.Request) {
	actionRequest := new(ActionRequest)

	err := json.NewDecoder(r.Body).Decode(actionRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	actionResponse, err := RunAction(actionRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := json.Marshal(actionResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseData)
}

type ActionRunner func (*ActionRequest) (*ActionResponse, error)

var actionRunnerFactory = make(map[string]ActionRunner)

func Register(name string, runner ActionRunner) {
	if runner == nil {
		logrus.Fatalf("Datastore factory %s does not exist.", name)
	}
	_, registered := actionRunnerFactory[name]
	if registered {
		logrus.Fatalf("Datastore factory %s already registered. Ignoring.", name)
	}
	actionRunnerFactory[name] = runner
}

func RunAction(req *ActionRequest) (*ActionResponse, error) {
	runner := actionRunnerFactory[req.Name]

	if runner != nil {
		return runner(req)
	} else {
		return nil, errors.Errorf("No handler for request: %v", req.Name)
	}
}