package main

func InitDownloadActions() {
	Register("potential_downloads", PotentialDownloads)
}

func PotentialDownloads(req *ActionRequest, cfg map[string]string) (*ActionResponse, error) {
	newCtx := make(map[string]string)

	showname, err := req.Entities.FirstEntityValue("showname")

	if err != nil {
		newCtx["missing_showname"] = "omitted"
	}

	showtype, err := req.Entities.FirstEntityValue("showtype")


	if err != nil {
		newCtx["missing_showtype"] = "omitted"
	} else {
		switch showtype {
		case "TV":
			newCtx["showtype"] = "TV"
			AddPotentialTVDownloads(showname, newCtx)
		case "Movie":
			newCtx["showtype"] = "movie"
			AddPotentialMovieDownloads(showname, newCtx)
		}
	}

	return &ActionResponse{
		Context: newCtx,
	}, nil
}

func AddPotentialTVDownloads(showname string, ctx map[string]string){
	ctx["showname"] = "my movie!"
	ctx["singleshowoption"] = "true"
}

func AddPotentialMovieDownloads(showname string, ctx map[string]string){
	ctx["showname"] = "my tv show!"
	ctx["singleshowoption"] = "true"
}