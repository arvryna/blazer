package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/arvryna/blazer/internals/chunk"
	"github.com/arvryna/blazer/internals/util"
)

func acceptedStatusCodes(code int) bool {
	table := map[int]string{
		200: "OK",
		206: "Partial content",
	}
	return table[code] != ""
}

// ConcurrentDownloader: Concurrently download the resource with specified concurrency value.
// it returns download status as bool
func ConcurrentDownloader(meta *FileMeta, thread int, sessionID string) (*chunk.Chunks, bool) {
	isDownloadComplete := true
	chunks := chunk.Chunks{Count: thread, TotalSize: int(meta.ContentLength)}
	chunks.ComputeChunks()
	var wg sync.WaitGroup
	for i, segment := range chunks.Segments {

		// if segment exist skip current segment download.
		if util.FileExists(util.SegmentFilePath(sessionID, i)) {
			continue
		}
		request, err := BuildRequest(http.MethodGet, meta.FileURL)
		if err != nil {
			fmt.Println("Could not build request: ", err)
			isDownloadComplete = false
		}
		wg.Add(1)

		i := i
		segment := segment

		// Downoading individual segments in go routines
		go func() {
			defer wg.Done()
			err = DownloadSegment(request, i, segment, sessionID)
			if err != nil {
				isDownloadComplete = false
				fmt.Println("Download segment failed: ", i, err)
			}
		}()
	}
	wg.Wait()
	return &chunks, isDownloadComplete
}

// DownloadSegment: download a specific piece of the bytes of the file that we want to download.
func DownloadSegment(request *http.Request, segmentID int, r chunk.Range, sessionID string) error {
	request.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r.Start, r.End))
	resp, err := HTTPClient().Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if !acceptedStatusCodes(resp.StatusCode) {
		return fmt.Errorf("received [un-expected] status code: %v resp: %v", resp.StatusCode, resp)
	}

	// read this byte by byte so you can show progress.
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(util.SegmentFilePath(sessionID, segmentID), bytes, os.ModePerm)
	if err != nil {
		return err
	}

	// Verify if segment is downloaded successfully.
	if len(bytes) == int(resp.ContentLength) {
		fmt.Println("Downloaded segment: ", segmentID)
	} else {
		return fmt.Errorf("incomplete segment: %v content-len %v", segmentID, resp.ContentLength)
	}
	return nil
}
