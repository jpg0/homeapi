package main

import (
	"net/http"
	"github.com/juju/errors"
	"net"
	"github.com/Sirupsen/logrus"
)

func restart(ac *GenericContext, cfg map[string]string) (*ActionResponse, error) {

	system, has_system := ac.MergeNew("system")

	if !has_system {
		ac.AddKey("missing_system")
		return ac.Response(), nil
	}

	var restartable Restartable

	switch system {
	case "kodi":
		restartable = NewRemoteRebootable("http://ubox:8808/")
	default:
		ac.AddKey("unknown_system")
		return ac.Response(), nil
	}

	restarting, err := restartable.Restart()

	if err != nil {
		return nil, errors.Annotate(err, "Failed to restart")
	}

	ac.Add("restarting", restarting)

	return ac.Response(), nil
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
	RegisterHandler("restart", AsGeneric(restart))
}