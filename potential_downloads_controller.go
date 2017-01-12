package main

import (
	"github.com/juju/errors"
	"fmt"
)

func InitPotentialDownloads() {
	RegisterHandler("potential_downloads", &PotentialDownloadsController{
		tvl: &SonarrAPIClientFake{
			tvdbid: 12345,
		},
	})
}

type PotentialDownloadsController struct {
	tvl TVLookup
}

func (dcm *PotentialDownloadsController) LookupShows(dc *PotentialDownloadsModel, cfg map[string]string) error {
	switch dc.ShowType {
	case TV:
		return dcm.tvl.LookupTVShows(dc, cfg)
	case Movie:
		panic(nil)
	default:
		return errors.Errorf("Unknown show type %v", dc.ShowType)
	}
}

func (dcm *PotentialDownloadsController) Run(req *APIAIRequest, config map[string]string) (*APIAIResponse, error) {

	res := NewAPIAIResponse("")

	dc := dcm.ToModel(req)

	if dc.Showquery != "" && dc.ShowType != "" {
		err := dcm.LookupShows(dc, config)

		if err != nil {
			return nil, errors.Annotate(err, "Failed to lookup shows")
		}

		if dc.Tvdbid != 0 {
			newCtx := NewAPIAIContext("show", 5)
			newCtx.Parameters["tvdbid"] = fmt.Sprint(dc.Tvdbid)
			newCtx.Parameters["showtype"] = string(dc.ShowType)

			res.AddContext(newCtx)

			res.SetMessage(fmt.Sprintf("Confirm download %v?", dc.ShowOptions[0]))

		} else { //multiple options

		}


		return res, nil

	}

	return nil, errors.New("Incomplete data")
}

func (dcm *PotentialDownloadsController) ToModel(req *APIAIRequest) *PotentialDownloadsModel {
	return &PotentialDownloadsModel{
		ShowType: showtype(req.Result.Parameters["showtype"]),
		Showquery: req.Result.Parameters["showquery"],
	}
}

type TVLookup interface {
	LookupTVShows(dc *PotentialDownloadsModel, cfg map[string]string) error
}


