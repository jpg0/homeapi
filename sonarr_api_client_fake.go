package main

type SonarrAPIClientFake struct {
	tvdbid int
	shows []string
}

func (sac *SonarrAPIClientFake) LookupTVShows(dc *PotentialDownloadsModel, cfg map[string]string) error {

	if sac.tvdbid != 0 {
		dc.FoundShow("found_by_fake_client", sac.tvdbid)
	} else {
		dc.FoundShows(sac.shows)
	}

	return nil
}

func (sac *SonarrAPIClientFake) DownloadTVShow(dc *DownloadingModel, cfg map[string]string) error {
	dc.Downloading = "true"
	dc.PosterUrl = "http://poster"

	return nil
}