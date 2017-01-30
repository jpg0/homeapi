package main

import (
	"net/http"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/juju/errors"
	"github.com/davecgh/go-spew/spew"
)

func ConfigureHandleAction(cfg *Configuration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		logrus.Debugf("Accepted new request")
		apiReq := &APIAIRequest{}

		err := json.NewDecoder(r.Body).Decode(apiReq)

		if err != nil {
			logrus.Errorf("Failed to unmarshal data: %v", err)
			actionError(err, w)
			return
		}

		actionResponse, err := RunAction(apiReq, cfg)

		if err != nil {
			logrus.Errorf("Failed to run action: %v", err)
			actionError(err, w)
			return
		}

		responseData, err := json.Marshal(actionResponse)

		if err != nil {
			logrus.Errorf("Failed to marshal data: %v", err)
			actionError(err, w)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.Write(responseData)
	}
}

func actionError(e error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	responseData, err := json.Marshal(NewAPIAIResponse(e.Error()))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(responseData)
}


func RunAction(req *APIAIRequest, cfg *Configuration) (*APIAIResponse, error) {
	logrus.Infof("Running action: %v", req.Result.Action)
	logrus.Debugf("Request details: %v", spew.Sprint(req))

	handler := GetHandler(req.Result.Action)

	if handler != nil {
		return handler.Run(req, cfg)
	} else {
		return nil, errors.Errorf("No handler for request: %v", req.Result.Action)
	}
}

func setupHandlers() {
	InitRestartActions();
	InitPotentialDownloads()
	InitDownloading()
	InitSelectShows()
	InitPhotosActions();
	InitShowDownloadsController();
}

