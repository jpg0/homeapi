package main

import (
	"path/filepath"
	"github.com/juju/errors"
	"fmt"
)

func InitPhotosActions() {
	Register("photo_uploads", PhotoUploads)
}

func PhotoUploads(ac *ActionContext, cfg map[string]string) (*ActionResponse, error) {

	files, err := filepath.Glob(cfg["photos_path"])

	if err != nil {
		return nil, errors.Annotatef(err, "Failed to list files for %v", "photos_path")
	}

	ac.Add("awaiting_upload", fmt.Sprintf("%v", len(files)))

	return ac.Response(), nil
}