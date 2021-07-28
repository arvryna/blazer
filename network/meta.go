package network

import (
	"fmt"
	"net/http"
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
		fmt.Println("Error downloading meta details of URL, ABORT")
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error downloading meta details of URL, ABORT")
	}
	meta := FileMeta{}
	meta.ContentLength = float64(resp.ContentLength)
	meta.ContentType = r.Header.Get("Content-Type")
	meta.FileUrl = url
	return &meta
}
