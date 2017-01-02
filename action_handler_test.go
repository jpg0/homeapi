package main

import (
	"testing"
	"github.com/Sirupsen/logrus"
)

func TestWithHydration(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	callThru := func(dc *DownloadContext, newCtx map[string]string, cfg map[string]string) (*ActionResponse, error) {
		return nil, nil
	}

	input := make(map[string]string)

	input["showname"] = "test_showname"
	input["tvdbid"] = "12345"
	input["showtype"] = "tv"
	input["showoptions"] = "abc&def&ghi"

	ar := WithHydration(callThru)

	ar(input, nil, nil)
	//
	//dc := &DownloadContext{}
	//
	//Hydrate(input, dc)
	//
	//assert.Equal(t, "test_showname", dc.Showquery)
	//assert.Equal(t, 12345, dc.Tvdbid)
	//assert.Equal(t, TV, dc.ShowType)
	//assert.Equal(t, []string{"abc", "def", "ghi"}, dc.ShowOptions)
}