package main


type DownloadingModel struct {
	Showquery string `ctx:"showquery"`
	Tvdbid int `ctx:"tvdbid"`
	PosterUrl string `ctx:"posterurl"`
	ShowType showtype `ctx:"showtype"`
	Downloading string
}