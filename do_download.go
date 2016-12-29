package main

import (
	"strings"
	"github.com/Sirupsen/logrus"
	"github.com/juju/errors"
	"github.com/jpg0/go-sonarr"
	"strconv"
)

func DoDownload(ac *ActionContext, cfg map[string]string) (*ActionResponse, error) {

	showtype, present := ac.MergeNew("showtype")

	if present {
		//has a show type
		switch strings.ToLower(showtype) {
		case "tv":
			tvdbid, err := strconv.ParseInt(ac.oldCtx["tvdbid"], 10, 0)

			if err != nil {
				return nil, errors.Annotatef(err, "Failed to parse tvdbid: %v", ac.oldCtx["tvdbid"])
			}

			downloading, err := DownloadTV(int(tvdbid), cfg)

			if err != nil {
				return nil, errors.Annotate(err, "Failed to download TV show")
			}

			//remove all as end of story
			ac.RemoveAllNow()
			ac.Add("downloading", downloading)

		case "movie":
			panic("not implemented")
		default:
			logrus.Infof("Unknown show type %v", showtype)
			ac.Remove("showtype")
			ac.Add("missing_showtype", "true")
		}
	} else {
		ac.Add("missing_showtype", "true")
	}

	return ac.Response(), nil
}

func DownloadTV(tvdbid int, cfg map[string]string) (string, error) {
	logrus.Debugf("Connecting to Sonarr at url: %v", cfg["sonarr_address"])

	sc, err := go_sonarr.NewSonarrClient(cfg["sonarr_address"], cfg["sonarr_apikey"])

	if err != nil {
		return "", errors.Annotate(err, "Failed to create Sonarr client")
	}

	spr, err := sc.CreateSeries(tvdbid)

	if err != nil {
		return "", errors.Annotate(err, "Failed to call Sonarr")
	}

	return getResponseText(spr), nil
}

func getResponseText(series go_sonarr.SonarrSeries) string {
	if series.Images != nil {
		for _, img := range series.Images {
			if img.CoverType == "poster" {
				return img.URL
			}
		}
	}

	return series.Title
}