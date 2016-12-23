package main

import (
	"net/http"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/juju/errors"
	"github.com/davecgh/go-spew/spew"
)

func ConfigureHandleAction(cfg map[string]string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		actionRequest := new(ActionRequest)

		err := json.NewDecoder(r.Body).Decode(actionRequest)

		if err != nil {
			actionError(err, w)
			return
		}

		actionResponse, err := RunAction(actionRequest, cfg)

		if err != nil {
			logrus.Errorf("Failed to run action: %v", err)
			actionError(err, w)
			return
		}

		responseData, err := json.Marshal(actionResponse)

		if err != nil {
			logrus.Errorf("Failed to mashal data: %v", err)
			actionError(err, w)
			return
		}

		w.Write(responseData)
	}
}

func actionError(e error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	responseData, err := json.Marshal(&ActionResponse{E:e})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseData)
}

type ActionRunner func(*ActionRequest, map[string]string) (*ActionResponse, error)

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

func RunAction(req *ActionRequest, cfg map[string]string) (*ActionResponse, error) {
	logrus.Infof("Running action: %v", req.Name)
	logrus.Debugf("Request details: %v", spew.Sprint(req))

	runner := actionRunnerFactory[req.Name]

	if runner != nil {
		return runner(req, cfg)
	} else {
		return nil, errors.Errorf("No handler for request: %v", req.Name)
	}
}

func setupHandlers() {
	InitListActions();
	InitRestartActions();
}

