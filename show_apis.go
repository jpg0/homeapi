package main

type Show struct {
	title string
	showid interface{}
}

type TVLookup interface {
	LookupTVShows(dc *PotentialDownloadsModel, cfg *Configuration) error
}

type MovieLookup interface {
	LookupMovies(dc *PotentialDownloadsModel, cfg *Configuration) error
}

type TVDownload interface {
	DownloadTVShow(dc *DownloadingModel, cfg *Configuration) error
}

type MovieDownload interface {
	DownloadMovie(dc *DownloadingModel, cfg *Configuration) error
}