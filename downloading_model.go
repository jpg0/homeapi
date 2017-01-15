package main


type DownloadingModel struct {
	Showname string `ctx:"showname"`
	Tvdbid int `ctx:"tvdbid"`
	PosterUrl string `ctx:"posterurl"`
	ShowType showtype `ctx:"showtype"`
	Downloading string
}