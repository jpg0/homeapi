package main

import (
	"io/ioutil"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/juju/errors"
	"path/filepath"
)

func loadConfiguration(path string) (*Configuration, error) {
	rv := &Configuration{}
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to read config file")
	}

	err = json.Unmarshal(file, &rv)

	if err != nil {
		return nil, errors.Annotate(err, "Failed to unmarshal config")
	}

	logrus.Debugf("Loading config from file: %v", path)

	rv.loadPath = path

	return rv, nil
}

type Configuration struct {
	KodiAddress string `json:"kodi_address"`
	SabnzbdAddress string `json:"sabnzbd_address"`
	SabnzbdApikey string `json:"sabnzbd_apikey"`
	SonarrAddress string `json:"sonarr_address"`
	SonarrApikey string `json:"sonarr_apikey"`
	CouchpotatoAddress string `json:"couchpotato_address"`
	CouchpotatoApikey string `json:"couchpotato_apikey"`
	HTTPPort string `json:"http_port"`
	HTTPSPort string `json:"https_port"`
	TLSCrtFile string `json:"tls_crt_file"`
	TLSKeyFile string `json:"tls_key_file"`
	PhotosPath string `json:"photos_path"`
	AuthPassword string `json:"auth_password"`

	loadPath string //source of the config file
}

func (c *Configuration) Resolve(path string) string {
	if !filepath.IsAbs(path) {
		return filepath.Join(filepath.Dir(c.loadPath), path)
	} else {
		return path
	}
}

