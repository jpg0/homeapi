package main

type Show struct {
	title string
	showid interface{}
}

type TVLookup interface {
	LookupTVShows(dc *PotentialDownloadsModel, cfg map[string]string) error
}

type MovieLookup interface {
	LookupMovies(dc *PotentialDownloadsModel, cfg map[string]string) error
}

type TVDownload interface {
	DownloadTVShow(dc *DownloadingModel, cfg map[string]string) error
}

type MovieDownload interface {
	DownloadMovie(dc *DownloadingModel, cfg map[string]string) error
}