package main

import (
	"strings"
	"github.com/Sirupsen/logrus"
)

func InitDownloadActions() {
	Register("potential_downloads", PotentialDownloads)
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
			AddPotentialTVDownloads(showname, ac)
		case "movie":
			ac.Add("showtype", "movie")
			AddPotentialMovieDownloads(showname, ac)
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

func AddPotentialTVDownloads(showname string, ac *ActionContext) {
	ac.Add("showname", "my movie!")
	ac.Add("singleshowoption", "true")
}

func AddPotentialMovieDownloads(showname string, ac *ActionContext) {
	ac.Add("showname", "my tv show!")
	ac.Add("singleshowoption", "true")
}