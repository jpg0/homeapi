package main


type DownloadingModel struct {
	Showname string `ctx:"showname"`
	Showid string `ctx:"showid"`
	PosterUrl string `ctx:"posterurl"`
	ShowType showtype `ctx:"showtype"`
	Downloading string
}