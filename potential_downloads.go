package main

import (
	"strings"
	"github.com/Sirupsen/logrus"
	"github.com/jpg0/go-sonarr"
	"github.com/juju/errors"
	"fmt"
)

func InitDownloadActions() {
	Register("potential_downloads", PotentialDownloads)
	Register("do_download", DoDownload)
}

func PotentialDownloads(ac *ActionContext, cfg map[string]string) (*ActionResponse, error) {

	showname, present := ac.MergeNew("showname")

	if present { //has a show name
		ac.Remove("missing_showname")
	} else {
		ac.Add("missing_showname", "true")
	}

	showtype, present := ac.MergeNew("showtype")

	if present { //has a show type
		switch strings.ToLower(showtype) {
		case "tv":
			ac.Add("showtype", "tv")
			err := AddPotentialTVDownloads(showname, ac, cfg)
			ac.Remove("missing_showtype")

			if err != nil {
				return nil, errors.Annotate(err, "Failed to lookup TV shows")
			}
		case "movie":
			ac.Add("showtype", "movie")
			err := AddPotentialMovieDownloads(showname, ac, cfg)
			ac.Remove("missing_showtype")

			if err != nil {
				return nil, errors.Annotate(err, "Failed to lookup movies")
			}
		default:
			logrus.Infof("Unknown show type %v", showtype)
			ac.Remove("showtype")
			ac.Add("missing_showtype", "true")
		}
	} else {
		ac.Add("missing_showtype", "true")
	}

	return ac.WriteTo(&ActionResponse{}), nil
}

func AddPotentialTVDownloads(showname string, ac *ActionContext, cfg map[string]string) error {
	ac.Remove("singleshowoption")
	ac.Remove("multipleshowoption")

	logrus.Debugf("Connecting to Sonarr at url: %v", cfg["sonarr_address"])

	sc, err := go_sonarr.NewSonarrClient(cfg["sonarr_address"], cfg["sonarr_apikey"])

	if err != nil {
		return errors.Annotate(err, "Failed to create Sonarr client")
	}

	slr, err := sc.SeriesLookup(showname)

	if err != nil {
		return errors.Annotate(err, "Failed to call Sonarr")
	}

	if len(slr) == 1 {
		ac.Add("singleshowoption", slr[0].Title)
		ac.Add("tvdbId", fmt.Sprintf("%v", slr[0].TvdbID))
	} else {

		shows := ""
		showIds := ""

		for _, show := range *slr {
			shows += show.Title + "\n"
			showIds += fmt.Sprintf("%v|", show.TvdbID)
		}


		ac.Add("multipleshowoption", shows)
		ac.Add("tvdbIds", showIds)
		ac.Remove("showname")
	}

	return nil
}

func AddPotentialMovieDownloads(showname string, ac *ActionContext, cfg map[string]string) error {
	ac.Add("singleshowoption", "my movie!")

	return nil
}