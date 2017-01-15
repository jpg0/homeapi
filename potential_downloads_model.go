package main

import "github.com/Sirupsen/logrus"

type showtype string

const (
	TV showtype = "tv"
	Movie showtype = "movie"
)

type PotentialDownloadsModel struct {
	ShowType showtype `ctx:"showtype"`
	Showquery string `ctx:"showquery"`
	ShowOptions []TVShow `ctx:"showoptions"`
}

func (dc *PotentialDownloadsModel) SetShowQuery(showquery string, cfg map[string]string) {
	if showquery == "" {
		panic("Showname not specified")
	}

	logrus.Debugf("Setting showquery to %v", showquery)

	dc.Showquery = showquery

}

func (dc *PotentialDownloadsModel) SetShowType(st showtype, cfg map[string]string) {
	if st == "" {
		panic("showtype is nil")
	}

	logrus.Debugf("Setting showtype to %v", st)

	dc.ShowType = st
}

func (dc *PotentialDownloadsModel) FoundShows(shows []TVShow) {
	dc.ShowOptions = shows
}