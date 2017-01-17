package main

import (
	"github.com/juju/errors"
	"fmt"
)

func InitDownloading() {
	RegisterHandler("do_download", &DownloadingController{
		//tdl: &SonarrAPIClientFake{},
		tdl: &SonarrAPIClient{},
		mdl: &CouchpotatoAPIClient{},
	})
}

func (dc *DownloadingController) Run(req *APIAIRequest, cfg *Configuration) (*APIAIResponse, error) {
	dm, err := dc.ToModel(req)

	if err != nil {
		return nil, errors.Annotatef(err, "Failed to read request as download request")
	}

	switch req.Result.Parameters["confirmation"] {
	case "no":
		return NewAPIAIResponse("Ok, not downloading anything"), nil
	case "yes":
		switch dm.ShowType {
		case TV:
			err = dc.tdl.DownloadTVShow(dm, cfg)

			if err != nil {
				return nil, errors.Annotatef(err, "Failed to download tv: %v", dm.Showid)
			}

			return NewAPIAIResponse(fmt.Sprintf("Downloading tv show: %v", dm.Showname)), nil
		case Movie:
			err = dc.mdl.DownloadMovie(dm, cfg)

			if err != nil {
				return nil, errors.Annotatef(err, "Failed to download movie: %v", dm.Showid)
			}

			return NewAPIAIResponse(fmt.Sprintf("Downloading movie: %v", dm.Showname)), nil
		default:
			return nil, errors.Errorf("Unknown showtype: %v", req.Result.Parameters["showtype"])
		}
	default:
		return nil, errors.Errorf("Unknown confirmation type: %v", req.Result.Parameters["confirmation"])
	}
}

func (dc *DownloadingController) ToModel(req *APIAIRequest) (*DownloadingModel, error) {
	dm := &DownloadingModel{}

	showtypeStr, present := req.Result.Parameters["showtype"]

	if present {
		dm.ShowType = showtype(showtypeStr)
	}

	showname, present := req.Result.Parameters["showname"]

	if present {
		dm.Showname = showname
	}

	showid, present := req.Result.Parameters["showid"]

	if present {
		dm.Showid = showid
	}

	posterurl, present := req.Result.Parameters["posterurl"]

	if present {
		dm.PosterUrl = posterurl
	}

	downloading, present := req.Result.Parameters["downloading"]

	if present {
		dm.Downloading = downloading
	}

	return dm, nil
}

type DownloadingController struct {
	tdl TVDownload
	mdl MovieDownload
}

func (dc *DownloadingController) Marshal(dm *DownloadingModel) map[string]string {
	ctx := make(map[string]string)

	ctx["showtype"] = string(dm.ShowType)
	ctx["showname"] = string(dm.Showname)

	if dm.Showid != "" {
		ctx["showid"] = dm.Showid
	}

	ctx["posterurl"] = dm.PosterUrl
	ctx["downloading"] = dm.Downloading

	return ctx
}