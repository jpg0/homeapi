package main

import (
	"github.com/juju/errors"
	"fmt"
	"strconv"
)

func InitDownloading() {
	RegisterHandler("do_download", &DownloadingController{
		//tdl: &SonarrAPIClientFake{},
		tdl: &SonarrAPIClient{},
	})
}

func (dc *DownloadingController) Run(req *APIAIRequest, cfg map[string]string) (*APIAIResponse, error) {
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
				return nil, errors.Annotatef(err, "Failed to download tv: %v", dm.Tvdbid)
			}

			return NewAPIAIResponse(fmt.Sprintf("Downloading: %v", dm.Showname)), nil
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

	tvdbid, present := req.Result.Parameters["tvdbid"]

	if present {
		tvdbidI64, err := strconv.ParseInt(tvdbid, 10, 0)

		if err != nil {
			return nil, errors.Annotatef(err, "Failed to parse tvdbid as int")
		}

		dm.Tvdbid = int(tvdbidI64)
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
}

func (dc *DownloadingController) Marshal(dm *DownloadingModel) map[string]string {
	ctx := make(map[string]string)

	ctx["showtype"] = string(dm.ShowType)
	ctx["showname"] = string(dm.Showname)

	if dm.Tvdbid != 0 {
		ctx["tvdbid"] = fmt.Sprint(dm.Tvdbid)
	}

	ctx["posterurl"] = dm.PosterUrl
	ctx["downloading"] = dm.Downloading

	return ctx
}

type TVDownload interface {
	DownloadTVShow(dc *DownloadingModel, cfg map[string]string) error
}