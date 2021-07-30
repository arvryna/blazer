package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/arvpyrna/blazer/data"
)

// merge output files concurrently using go channel or something
func ConcurrentDownloader(meta *FileMeta, thread int) {
	fmt.Println("Initiating download... dispatching workers")
	chunks := data.CalculateChunks(int(meta.ContentLength), thread)
	var wg sync.WaitGroup
	for i, segment := range chunks.Segments {
		request, err := BuildRequest(http.MethodGet, meta.FileUrl)
		if err != nil {
			fmt.Println(err)
		}
		// start before concurrency
		wg.Add(1)
		// capturing values as they change
		i := i
		segment := segment
		go func() {
			defer wg.Done()
			DownloadSegment(request, i, segment)
		}()
	}
	wg.Wait()
	data.MergeFiles(chunks, meta.FileName)
}

func DownloadSegment(request *http.Request, i int, r data.Range) {
	request.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r.Start, r.End))
	resp, err := HTTPClient().Do(request)
	if err != nil {
		fmt.Println(err)
	}

	// handle error
	// read this byte by byte so you can show progress
	//TODO: Check if resp is nil, also check error codes
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(data.SegmentFilePath(i), bytes, os.ModePerm)
	// check if bytes written is same as content size
	fmt.Println("Downloaded segment: ", i)
}
