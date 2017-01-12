package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/jpg0/go-sonarr"
	"github.com/juju/errors"
	"net/url"
)

type SonarrAPIClient struct {

}

func (sac *SonarrAPIClient) LookupTVShows(dc *PotentialDownloadsModel, cfg map[string]string) error {
	logrus.Debugf("Connecting to Sonarr at url: %v", cfg["sonarr_address"])

	sc, err := go_sonarr.NewSonarrClient(cfg["sonarr_address"], cfg["sonarr_apikey"])

	if err != nil {
		return errors.Annotate(err, "Failed to create Sonarr client")
	}

	slr, err := sc.SeriesLookup(dc.Showquery)

	if err != nil {
		return errors.Annotate(err, "Failed to call Sonarr")
	}

	if len (slr) == 0 {
		dc.FoundShows([]string{})
	} else if len(slr) == 1 {
		dc.FoundShow(slr[0].Title, slr[0].TvdbID)
	} else {
		showoptions := make([]string, len(slr))

		for _, show := range slr {
			showoptions = append(showoptions, show.Title)
		}

		dc.FoundShows(showoptions)
	}

	return nil
}

func (sac *SonarrAPIClient) DownloadTVShow(dc *DownloadingModel, cfg map[string]string) error {
	logrus.Debugf("Connecting to Sonarr at url: %v", cfg["sonarr_address"])

	sc, err := go_sonarr.NewSonarrClient(cfg["sonarr_address"], cfg["sonarr_apikey"])

	if err != nil {
		return errors.Annotate(err, "Failed to create Sonarr client")
	}

	spr, err := sc.CreateSeries(dc.Tvdbid)

	if err != nil {
		return errors.Annotate(err, "Failed to call Sonarr")
	}

	dc.PosterUrl = posterUrl(spr, cfg["sonarr_address"])
	dc.Downloading = "true"

	return nil
}

func posterUrl(series go_sonarr.SonarrSeries, sonarr_address string) string {
	if series.Images != nil {
		for _, img := range series.Images {
			if img.CoverType == "poster" {
				// download and convert to data URI

				root, _ := url.Parse(sonarr_address)
				imgUrl, _ := url.Parse(img.URL)

				dataURI, err := toDataURI(root.ResolveReference(imgUrl))

				if err != nil {
					logrus.Debugf("Failed to get image URL: %v", err)
					break
				} else {
					return dataURI
				}
			}
		}
	}

	return ""
}
