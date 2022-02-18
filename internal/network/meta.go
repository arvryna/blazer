package network

import (
	"fmt"
	"net/http"
	"path"
)

type FileMeta struct {
	FileURL       string
	FileName      string
	ContentLength float64
	ServerName    string
	Age           int
	ContentType   string
	AcceptRanges  string
}

// Fetch the meta details of a given URL.
func (m *FileMeta) Fetch(url string) error {
	r, err := BuildRequest(http.MethodHead, url)
	if err != nil {
		return err
	}

	resp, err := HTTPClient().Do(r)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Rreceived un-expected status code: %v resp: %v", resp.StatusCode, resp)
	}

	m.FileURL = url
	m.ContentLength = float64(resp.ContentLength)
	m.ContentType = r.Header.Get("Content-Type")
	m.FileName = path.Base(r.URL.Path)
	m.AcceptRanges = resp.Header.Get("Accept-Ranges")
	return nil
}
