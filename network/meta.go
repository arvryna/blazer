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

// check for null response and in such cases you can use, default thread to download
func GetFileMeta(url string) *FileMeta {
	r, err := BuildRequest(http.MethodHead, url)
	if err != nil {
		fmt.Printf("Error building URL %v", err)
	}

	resp, err := HTTPClient().Do(r)
	if err != nil {
		fmt.Printf("Error fetching meta details of URL %v", err)
	}

	meta := FileMeta{
		FileUrl:       url,
		ContentLength: float64(resp.ContentLength),
		ContentType:   r.Header.Get("Content-Type"),
		FileName:      path.Base(r.URL.Path),
	}
	return &meta
}
