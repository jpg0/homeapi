package main

import (
	"path/filepath"
	"github.com/juju/errors"
	"fmt"
)

const DEFAULT_PHOTOS_PATH = "/photos/*.*"

func InitPhotosActions() {
	Register("photo_uploads", PhotoUploads)
}

func PhotoUploads(ac *ActionContext, cfg map[string]string) (*ActionResponse, error) {

	var path string
	var present bool
	if path, present = cfg["photos_path"] ; !present {
		path = DEFAULT_PHOTOS_PATH
	}

	files, err := filepath.Glob(path)

	if err != nil {
		return nil, errors.Annotatef(err, "Failed to list files for %v", path)
	}

	ac.Add("awaiting_upload", fmt.Sprintf("%v", len(files)))

	return ac.Response(), nil
}