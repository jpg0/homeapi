package main

import (
	"github.com/juju/errors"
	"fmt"
)

func InitPotentialDownloads() {
	RegisterHandler("potential_downloads", &PotentialDownloadsController{
		//tvl: &SonarrAPIClientFake{
		//	shows: []TVShow {
		//		{title:"fake show 1", tvdbid:12345},
		//		{title:"fake show 2", tvdbid:12346},
		//	},
		//},
		tvl: &SonarrAPIClient{},
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

		if len(dc.ShowOptions) == 0 {
			newCtx := NewAPIAIContext("no_show", 1)
			newCtx.Parameters["showtype"] = string(dc.ShowType)
			newCtx.Parameters["showname"] = dc.Showquery

			res.AddContext(newCtx)

			res.SetMessage(fmt.Sprintf("No show found for %v. Any others?", dc.Showquery))
		} else if len(dc.ShowOptions) == 1 {
			newCtx := NewAPIAIContext("show", 5)
			newCtx.Parameters["tvdbid"] = fmt.Sprint(dc.ShowOptions[0].tvdbid)
			newCtx.Parameters["showtype"] = string(dc.ShowType)
			newCtx.Parameters["showname"] = dc.ShowOptions[0].title

			res.AddContext(newCtx)

			res.SetMessage(fmt.Sprintf("Confirm download %v?", dc.ShowOptions[0].title))

		} else { //multiple options

			newCtx := NewAPIAIContext("show_options", 1)
			msg := "Please select from the following matches:\n"

			for i := range dc.ShowOptions {
				show := dc.ShowOptions[i]

				newCtx.Parameters[fmt.Sprint(i)] = map[string]string{
					"title":show.title,
					"tvdbid":fmt.Sprint(show.tvdbid),
					"showtype":string(dc.ShowType),
				}

				msg += fmt.Sprintf("%v) %v\n", i, show.title)
			}

			res.AddContext(newCtx)
			res.SetMessage(msg)
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


