package main

import (
	"github.com/tehjojo/go-sabnzbd"
	"github.com/juju/errors"
	"bytes"
	"fmt"
)

type ListActions Configuration

func list(req *ActionRequest, cfg *Configuration) (*ActionResponse, error) {

	s, err := sabnzbd.New(sabnzbd.Addr(cfg["sabnzbd_address"]), sabnzbd.ApikeyAuth(cfg["sabnzbd_apikey"]))

	if err != nil {
		return nil, errors.Annotate(err, "Failed to create NZB client")
	}

	history, err := s.History(0, 10)

	var buffer bytes.Buffer

	for _, slot := range history.Slots {
		fmt.Fprintf(buffer, "%v\n", slot)
	}

	period := "7 days"
	data := string(buffer)

	return &ActionResponse{
		Context: map[string]string{
			"period": period,
			"data": data,
		},
	}, nil
}


func InitListActions() {
	Register("list", list)
}