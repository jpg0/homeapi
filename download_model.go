package main

type showtype string

const (
	TV showtype = "tv"
	Movie showtype = "movie"
)

type DownloadingModel struct {
	Showname string `ctx:"showname"`
	Tvdbid int `ctx:"tvdbid"`
	PosterUrl string `ctx:"posterurl"`
}

type DownloadContext struct {
	ShowType showtype `ctx:"showtype"`
	Showquery string `ctx:"showquery"`
	Tvdbid int `ctx:"tvdbid"`
	ShowOptions []string `ctx:"showoptions"`
}

func (dc *DownloadContext) Apply(newCtx map[string]string, cfg map[string]string) {
	newShowType, present := newCtx["showtype"]

	if present {
		dc.SetShowType(showtype(newShowType), cfg)
	}

	newShowName, present := newCtx["showname"]

	if present {
		dc.SetShowName(newShowName, cfg)
	}
}

func (dc *DownloadContext) SetShowName(showname string, cfg map[string]string) {
	if showname == "" {
		panic("Showname not specified")
	}

	dc.Showquery = showname

}

func (dc *DownloadContext) SetShowType(st showtype, cfg map[string]string) {
	if st == "" {
		panic("showtype is nil")
	}

	dc.ShowType = st
}

func (dc *DownloadContext) FoundShow(showname string, tvdbid int) {

	dc.ShowOptions = []string{showname}
	dc.Tvdbid = tvdbid
}

func (dc *DownloadContext) FoundShows(shownames []string) {

	dc.Tvdbid = 0
	dc.ShowOptions = shownames
}