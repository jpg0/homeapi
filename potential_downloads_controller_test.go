package main
//
//import (
//	"testing"
//	"github.com/stretchr/testify/assert"
//)
//
//type ProxyTVClient struct {
//	proxy func(dc *PotentialDownloadsModel, cfg map[string]string) error
//}
//
//func (ptc *ProxyTVClient) LookupTVShows(dc *PotentialDownloadsModel, cfg map[string]string) error {
//	return ptc.proxy(dc, cfg)
//}
//
//func TestNoMatch(t *testing.T) {
//	dcm := &PotentialDownloadsController{
//		tvl: &ProxyTVClient{
//			proxy: func(dc *PotentialDownloadsModel, cfg map[string]string) error {
//				//noop
//				return nil
//			},
//		},
//	}
//
//	ar, err := dcm.Run(
//		map[string]string{
//			"showtype":"tv",
//			"showquery":"foo",
//		},
//		nil,
//		nil,
//	)
//
//	if err != nil {
//		t.Errorf("Failed to run: %v", err)
//	}
//
//	assert.Nil(t, ar.AddContext)
//	assert.Nil(t, ar.E)
//	assert.Empty(t, ar.Message)
//	assert.Nil(t, ar.RemoveContext)
//
//	c := ar.ReplaceContext
//
//	assert.Equal(t, c["showtype"], "tv")
//	assert.Equal(t, c["showquery"], "foo")
//	assert.Empty(t, c["tvdbid"])
//	assert.Empty(t, c["showoptions"])
//}
//
//func TestSingleMatch(t *testing.T) {
//	dcm := &PotentialDownloadsController{
//		tvl: &ProxyTVClient{
//			proxy: func(dc *PotentialDownloadsModel, cfg map[string]string) error {
//				//noop
//				dc.FoundShow("foo1", 123)
//
//				return nil
//			},
//		},
//	}
//
//	ar, err := dcm.Run(
//		map[string]string{
//			"showtype":"tv",
//			"showquery":"foo",
//		},
//		nil,
//		nil,
//	)
//
//	if err != nil {
//		t.Errorf("Failed to run: %v", err)
//	}
//
//	assert.Nil(t, ar.AddContext)
//	assert.Nil(t, ar.E)
//	assert.Empty(t, ar.Message)
//	assert.Nil(t, ar.RemoveContext)
//
//	c := ar.ReplaceContext
//
//	assert.Equal(t, c["showtype"], "tv")
//	assert.Equal(t, c["showquery"], "foo")
//	assert.Equal(t, c["tvdbid"], "123")
//	assert.Equal(t, c["showoptions"], "foo1")
//}
//
//func TestTwoMatches(t *testing.T) {
//	dcm := &PotentialDownloadsController{
//		tvl: &ProxyTVClient{
//			proxy: func(dc *PotentialDownloadsModel, cfg map[string]string) error {
//				//noop
//				dc.FoundShows([]string{"foo1", "foo2"})
//
//				return nil
//			},
//		},
//	}
//
//	ar, err := dcm.Run(
//		map[string]string{
//			"showtype":"tv",
//			"showquery":"foo",
//		},
//		nil,
//		nil,
//	)
//
//	if err != nil {
//		t.Errorf("Failed to run: %v", err)
//	}
//
//	assert.Nil(t, ar.AddContext)
//	assert.Nil(t, ar.E)
//	assert.Empty(t, ar.Message)
//	assert.Nil(t, ar.RemoveContext)
//
//	c := ar.ReplaceContext
//
//	assert.Equal(t, c["showtype"], "tv")
//	assert.Equal(t, c["showquery"], "foo")
//	assert.Empty(t, c["tvdbid"])
//	assert.Equal(t, c["showoptions"], "foo1&foo2")
//}