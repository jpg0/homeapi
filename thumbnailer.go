package main

import (
	"net/url"
	"github.com/Sirupsen/logrus"
	"github.com/nfnt/resize"
	"github.com/vincent-petithory/dataurl"
	"net/http"
	"github.com/juju/errors"
	"image"
	"image/jpeg"
	"image/png"
	"bytes"
	"bufio"
)

func toDataURI(imgUrl *url.URL) (string, error) {

	logrus.Debugf("Retrieving image from %v", imgUrl.String())

	res, err := http.DefaultClient.Get(imgUrl.String())

	if err != nil {
		return "", errors.Annotate(err, "Failed to download image from Sonarr")
	}

	ct := res.Header.Get("Content-Type")

	var img image.Image

	switch ct {
	case "image/jpeg":
		img, err = jpeg.Decode(res.Body)
	default:
		return "", errors.Errorf("Unknown content type for image: %v", ct)
	}

	if err != nil {
		return "", errors.Annotate(err, "Failed to parse image")
	}

	thumb := resize.Thumbnail(0, 200, img, resize.Bilinear)

	var b bytes.Buffer

	png.Encode(bufio.NewWriter(&b), thumb)

	return dataurl.EncodeBytes(b.Bytes()), nil
}

