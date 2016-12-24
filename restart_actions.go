package main

import (
	"net/http"
	"github.com/juju/errors"
	"net"
	"github.com/Sirupsen/logrus"
)

func restart(req *ActionRequest, cfg map[string]string) (*ActionResponse, error) {

	system, err := req.Entities.FirstEntityValue("system")

	if err != nil {
		return nil, errors.Errorf("No system specified to restart")
	}

	var restartable Restartable

	switch system {
	case "kodi":
		restartable = NewRemoteRebootable("http://ubox:8808/")
	}

	restarting, err := restartable.Restart()

	if err != nil {
		return nil, errors.Annotate(err, "Failed to restart")
	}

	return &ActionResponse{
		Context: map[string]string{"restarting": restarting},
	}, nil
}

type Restartable interface {
	Restart() (string, error)
}

func NewRemoteRebootable(address string) *RemoteRebootable {
	return &RemoteRebootable{address:address}
}

type RemoteRebootable struct {
	address string
}

func (r *RemoteRebootable) Restart() (string, error) {
	request, err := http.NewRequest("DELETE", r.address, nil)

	if (err != nil) {
		return "", errors.Annotate(err, "Failed to construct remote request")
	}

	response, err := http.DefaultClient.Do(request)

	if (err != nil) {
		return "", errors.Annotate(err, "Failed to invoke remote request")
	}

	if (response.StatusCode != http.StatusAccepted) {
		return "", errors.Errorf("Failed to invoke action, response code is: %v", response.StatusCode)
	}

	host, _, err := net.SplitHostPort(request.Host)

	if err != nil {
		logrus.Errorf("Failed to split host:port in %v: %v", request.Host, err)
		host = request.Host //just use the raw request.Host
	}

	return host, nil
}

func InitRestartActions() {
	Register("restart", restart)
}