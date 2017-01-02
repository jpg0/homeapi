package main

import (
	"github.com/tehjojo/go-sabnzbd"
	"github.com/juju/errors"
	"bytes"
	"fmt"
)

type ListActions Configuration

func list(ac *GenericContext, cfg map[string]string) (*ActionResponse, error) {

	s, err := sabnzbd.New(sabnzbd.Addr(cfg["sabnzbd_address"]), sabnzbd.ApikeyAuth(cfg["sabnzbd_apikey"]))

	if err != nil {
		return nil, errors.Annotate(err, "Failed to create NZB client")
	}

	history, err := s.History(0, 10)

	var buffer bytes.Buffer

	for _, slot := range history.Slots {
		buffer.WriteString(fmt.Sprintf("%v\n", slot.Name))
	}

	ac.Add("period", "7 days")
	ac.Add("data", buffer.String())

	return ac.Response(), nil
}


func InitListActions() {
	Register("list", AsGeneric(list))
}