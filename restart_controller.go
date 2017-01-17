package main


import (
	"net/http"
	"github.com/juju/errors"
	"net"
	"github.com/Sirupsen/logrus"
	"fmt"
)

type RestartController struct {

}

func (rc *RestartController) Run(req *APIAIRequest, config *Configuration) (*APIAIResponse, error) {

	logrus.Debugf("Restart system command received")

	system, has_system := req.Result.Parameters["system"]

	if !has_system {
		return nil, errors.New("'system' parameter not present")
	}

	var restartable Restartable

	switch system {
	case "ubox":
		restartable = NewRemoteRebootable("http://ubox:8808/")
	default:
		return nil, errors.Errorf("unknown system: %v", system)
	}

	restarting, err := restartable.Restart()

	if err != nil {
		return nil, errors.Annotate(err, "Failed to restart")
	}

	logrus.Infof("Sent restart request to: %v", restarting)
	return NewAPIAIResponse(fmt.Sprintf("Restarting %v", restarting)), nil
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
	RegisterHandler("restart", &RestartController{})
}