package main

func getDownloads(*ActionRequest) (*ActionResponse, error) {
	return &ActionResponse{
		Message: "some stuff",
	}, nil
}


func InitDownloadActions() {
	Register("getDownloads", getDownloads)
}