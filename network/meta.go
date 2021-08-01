package network

import (
	"fmt"
	"net/http"
	"path"
)

type FileMeta struct {
	FileUrl       string
	FileName      string
	ContentLength float64
	ServerName    string
	Age           int
	ContentType   string
}

func GetFileMeta(url string) (*FileMeta, error) {
	r, err := BuildRequest(http.MethodHead, url)
	if err != nil {
		return nil, err
	}

	resp, err := HTTPClient().Do(r)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("received un-expected status code: %v resp: %v", resp.StatusCode, resp)
	}

	if err != nil {
		return nil, err
	}

	meta := FileMeta{
		FileUrl:       url,
		ContentLength: float64(resp.ContentLength),
		ContentType:   r.Header.Get("Content-Type"),
		FileName:      path.Base(r.URL.Path),
	}
	return &meta, nil
}
