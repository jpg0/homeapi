package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/jpg0/go-couchpotato"
	"github.com/juju/errors"
)

type CouchpotatoAPIClient struct {

}

func (cac *CouchpotatoAPIClient) LookupMovies(dc *PotentialDownloadsModel, cfg map[string]string) error {
	logrus.Debugf("Connecting to Couchpotato at url: %v", cfg["sonarr_address"])

	cc, err := go_couchpotato.NewCouchpotatoClient(cfg["couchpotato_address"], cfg["couchpotato_apikey"])

	if err != nil {
		return errors.Annotate(err, "Failed to create Couchpotato client")
	}

	movies, err := cc.SearchMovies(dc.Showquery)

	if err != nil {
		return errors.Annotate(err, "Failed to call Couchpotato")
	}

	shows := make([]Show, len(movies))
	for i := range movies {
		shows[i] = Show{title:movies[i].OriginalTitle, showid:movies[i].Imdb}
	}
	dc.FoundShows(shows)

	return nil
}

func (cac *CouchpotatoAPIClient) DownloadMovie(dc *DownloadingModel, cfg map[string]string) error {
	logrus.Debugf("Connecting to Couchpotato at url: %v", cfg["sonarr_address"])

	cc, err := go_couchpotato.NewCouchpotatoClient(cfg["couchpotato_address"], cfg["couchpotato_apikey"])

	if err != nil {
		return errors.Annotate(err, "Failed to create Couchpotato client")
	}

	_, err = cc.AddMovie(dc.Showname, dc.Showid)

	if err != nil {
		return errors.Annotate(err, "Failed to call Couchpotato")
	}

	return nil
}