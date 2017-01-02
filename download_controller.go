package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/jpg0/go-sonarr"
	"github.com/juju/errors"
	"net/url"
	"io/ioutil"
	"github.com/vincent-petithory/dataurl"
	"net/http"
)

func InitDownloadActions() {
	Register("potential_downloads", WithHydration(PotentialDownloads))
	Register("do_download", WithHydration(DoDownload))
}

func PotentialDownloads(dc *DownloadContext, newCtx map[string]string, cfg map[string]string) (*ActionResponse, error) {
	dc.Apply(newCtx, cfg)

	err := LookupShows(dc, cfg)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to lookup shows")
	}

	return DehydratedResponse(dc), nil
}

func LookupShows(dc *DownloadContext, cfg map[string]string) error {
	switch dc.ShowType {
	case TV:
		return LookupTVShows(dc, cfg)
	case Movie:
		panic(nil)
	default:
		return errors.Errorf("Unknown show type %v", dc.ShowType)
	}
}

func LookupTVShows(dc *DownloadContext, cfg map[string]string) error {
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

func DoDownload(dc *DownloadContext, newCtx map[string]string, cfg map[string]string) (*ActionResponse, error) {
	if len(newCtx) > 0 {
		return nil, errors.New("New context data for download")
	}

	if dc.Tvdbid == 0 {
		return nil, errors.New("No tvdbid set to download show")
	}


	switch dc.ShowType {
	case TV:
		return DownloadTV(dc.Tvdbid, cfg)
	default:
		return nil, errors.Errorf("Unknown show type %v", dc.ShowType)
	}
}

func DownloadTV(tvdbid int, cfg map[string]string) (*ActionResponse, error) {
	logrus.Debugf("Connecting to Sonarr at url: %v", cfg["sonarr_address"])

	sc, err := go_sonarr.NewSonarrClient(cfg["sonarr_address"], cfg["sonarr_apikey"])

	if err != nil {
		return nil, errors.Annotate(err, "Failed to create Sonarr client")
	}

	spr, err := sc.CreateSeries(tvdbid)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to call Sonarr")
	}

	dm := &DownloadingModel{
		Showname: spr.Title,
		Tvdbid: tvdbid,
		PosterUrl: posterUrl(spr, cfg["sonarr_address"]),
	}

	return DehydratedResponse(dm), nil
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
					logrus.Debugf("Failed to ")
					break
				} else {
					return dataURI
				}
			}
		}
	}

	return ""
}

func toDataURI(imgUrl *url.URL) (string, error) {

	res, err := http.DefaultClient.Get(imgUrl.String())

	if err != nil {
		return "", errors.Annotate(err, "Failed to download image from Sonarr")
	}

	raw, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", errors.Annotate(err, "Failed to read image data from Sonarr")
	}

	return dataurl.EncodeBytes(raw), nil
}