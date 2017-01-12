package main

import (
	"testing"
	"net/url"
	"fmt"
	"github.com/Sirupsen/logrus"
)

func TestSonarr(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	u, e := url.Parse("http://vault:8989/MediaCover/161/poster.jpg")

	if e != nil {
		t.Fatal(e)
	}

	url, e := toDataURI(u)

	if e != nil {
		t.Fatal(e)
	}

	fmt.Printf("URL: ", url)
}