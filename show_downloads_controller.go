package main

import (
"github.com/tehjojo/go-sabnzbd"
"github.com/juju/errors"
"bytes"
"fmt"
)

type ShowDownloadsController struct {

}

func (sdc *ShowDownloadsController) Run(req *APIAIRequest, cfg *Configuration) (*APIAIResponse, error) {

	s, err := sabnzbd.New(sabnzbd.Addr(cfg.SabnzbdAddress), sabnzbd.ApikeyAuth(cfg.SabnzbdApikey))

	if err != nil {
		return nil, errors.Annotate(err, "Failed to create NZB client")
	}

	history, err := s.History(0, 10)

	var buffer bytes.Buffer

	for _, slot := range history.Slots {
		buffer.WriteString(fmt.Sprintf("%v\n", slot.Name))
	}

	//ac.Add("period", "7 days")
	//ac.Add("data", buffer.String())

	//return ac.Response(), nil
	return NewAPIAIResponse(buffer.String()), nil
}


func InitShowDownloadsController() {
	RegisterHandler("list_downloads", &ShowDownloadsController{})
}