package main

import (
	"path/filepath"
	"fmt"
	"github.com/juju/errors"
)

const DEFAULT_PHOTOS_PATH = "/photos/*.*"

func InitPhotosActions() {
	RegisterHandler("photo_uploads", &PhotosController{})
}

type PhotosController struct {

}

func (pc *PhotosController) Run(req *APIAIRequest, cfg map[string]string) (*APIAIResponse, error) {

	var path string
	var present bool
	if path, present = cfg["photos_path"] ; !present {
		path = DEFAULT_PHOTOS_PATH
	}

	files, err := filepath.Glob(path)

	if err != nil {
		return nil, errors.Annotatef(err, "Failed to list files for %v", path)
	}

	return NewAPIAIResponse(fmt.Sprintf("There are %v photos awaiting upload", len(files))), nil
}