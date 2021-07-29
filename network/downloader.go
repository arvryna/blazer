package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/arvpyrna/blazer/data"
)

func BuildRequest(method string, url string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, nil)
	r.Header.Set("User-Agent", "Blazer")
	return r, err
}

// fatal error: concurrent map iteration and map write
// merge output files concurrently using go channel or something
func ConcurrentDownloader(meta *FileMeta, thread int) {
	fmt.Println("Initiating download... dispatching workers")
	chunks := data.CalculateChunks(int(meta.ContentLength), thread)
	var wg sync.WaitGroup
	for i, segment := range chunks.Segments {
		request, _ := BuildRequest(http.MethodGet, meta.FileUrl)
		// start before concurrency
		wg.Add(1)
		// capturing values as they change
		i := i
		segment := segment
		go func() {
			defer wg.Done() // defer is good pattern than trying to close in a specific place
			DownloadSegment(request, i, segment)
		}()
	}
	wg.Wait()
	fmt.Println("Merging files..")
	mergefiles(chunks)
}

// To avoid TLS handshake
// https://stackoverflow.com/questions/41719797/tls-handshake-timeout-on-requesting-data-concurrently-from-api
func HTTPClient() *http.Client {
	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 600 * time.Second,
	}
	c := &http.Client{
		Transport: t,
	}
	return c
}

func DownloadSegment(request *http.Request, i int, r data.Range) {
	request.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r.Start, r.End))
	resp, err := HTTPClient().Do(request)
	if err != nil {
		fmt.Println(err)
	}

	// handle error
	// read this byte by byte so you can show progress
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(fmt.Sprintf("out/segment-%v.pdf", i), data, os.ModePerm)
	fmt.Println("Downloaded segment: ", i)
}

func mergefiles(chunks *data.Chunks) {
	f, _ := os.OpenFile("out/fin.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer f.Close()
	for i := range chunks.Segments {
		fileName := fmt.Sprintf("out/segment-%v.pdf", i)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		err = os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
		}
	}
}
