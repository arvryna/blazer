package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/arvpyrna/blazer/data"
)

func ConcurrentDownloader(meta *FileMeta, thread int, outputName string) {
	fmt.Println("Download the file in threads", thread)
	chunks := data.CalculateChunks(int(meta.ContentLength), thread)
	var wg sync.WaitGroup
	for i, segment := range chunks.Segments {
		// if segment exist skip current segment download
		if data.FileExists(data.SegmentFilePath(data.SESSION_ID, i)) {
			// fmt.Println("Segment Id: ", i, "already downloaded")
			continue
		}
		request, err := BuildRequest(http.MethodGet, meta.FileUrl)
		if err != nil {
			fmt.Println(err)
		}
		wg.Add(1)
		i := i
		segment := segment
		go func() {
			defer wg.Done()
			DownloadSegment(request, i, segment)
		}()
	}
	wg.Wait()
	data.MergeFiles(chunks, outputName)
}

func DownloadSegment(request *http.Request, i int, r data.Range) {
	request.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r.Start, r.End))
	resp, err := HTTPClient().Do(request)
	if err != nil {
		fmt.Println(err)
	}

	// read this byte by byte so you can show progress
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(data.SegmentFilePath(data.SESSION_ID, i), bytes, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	// Verify if segment is downloaded successfully
	if len(bytes) == int(resp.ContentLength) {
		fmt.Println("Downloaded segment: ", i)
	} else {
		// Push undownloaded segments to a channel for retry
		fmt.Println("Can't Download segment: ", i, "content-len", resp.ContentLength)
	}
}
