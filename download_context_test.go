package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHydrate(t *testing.T) {

	input := make(map[string]string)

	input["showname"] = "test_showname"
	input["tvdbid"] = "12345"
	input["showtype"] = "tv"
	input["showoptions"] = "abc&def&ghi"

	dc := &DownloadContext{}

	Hydrate(input, dc)

	assert.Equal(t, "test_showname", dc.Showname)
	assert.Equal(t, 12345, dc.Tvdbid)
	assert.Equal(t, TV, dc.ShowType)
	assert.Equal(t, []string{"abc", "def", "ghi"}, dc.ShowOptions)
}

func TestResponse(t *testing.T) {

	dc := &DownloadContext{
		Showname:"test_showname",
		Tvdbid:12345,
		ShowType:TV,
		ShowOptions:[]string{"abc", "def", "ghi"},
	}

	ctx := Dehydrate(dc)

	input := make(map[string]string)

	input["showname"] = "test_showname"
	input["tvdbid"] = "12345"
	input["showtype"] = "tv"
	input["showoptions"] = "abc&def&ghi"

	assert.Equal(t, "test_showname", ctx["showname"])
	assert.Equal(t, "12345", ctx["tvdbid"])
	assert.Equal(t, "tv", ctx["showtype"])
	assert.Equal(t, "abc&def&ghi", ctx["showoptions"])
}