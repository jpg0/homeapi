package main

type SonarrAPIClientFake struct {
	tvdbid int
	shows []Show
}

func (sac *SonarrAPIClientFake) LookupTVShows(dc *PotentialDownloadsModel, cfg map[string]string) error {

	dc.FoundShows(sac.shows)

	return nil
}

func (sac *SonarrAPIClientFake) DownloadTVShow(dc *DownloadingModel, cfg map[string]string) error {
	dc.Downloading = "true"
	dc.PosterUrl = "http://poster"

	return nil
}