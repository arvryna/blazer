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

type Dispatcher struct {
	ThreadCount int
	SessionID   string
	Meta        *FileMeta
}

// Concurrently download the resource with specified concurrency value.
// it returns download status as bool
func (d *Dispatcher) InitiateConcurrentDispatch() (*chunk.Chunks, bool) {
	isDownloadComplete := true

	chunks := chunk.Chunks{Count: d.ThreadCount, TotalSize: int(d.Meta.ContentLength)}
	chunks.ComputeChunks()

	var wg sync.WaitGroup
	wg.Add(d.ThreadCount)

	for i, segment := range chunks.Segments {

		// if segment exist skip current segment download.
		if util.FileExists(util.SegmentFilePath(d.SessionID, i)) {
			continue
		}

		request, err := BuildRequest(http.MethodGet, d.Meta.FileURL)

		if err != nil {
			fmt.Println("Could not build request: ", err)
			isDownloadComplete = false
		}

		i := i
		segment := segment

		// Downoading individual segments in go routines
		go func() {
			defer wg.Done()
			err = d.downloadSegment(request, i, segment)
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
func (d *Dispatcher) downloadSegment(request *http.Request, segmentID int, r chunk.Range) error {
	request.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r.Start, r.End))
	resp, err := HTTPClient().Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if !d.acceptedStatusCodes(resp.StatusCode) {
		return fmt.Errorf("received [un-expected] status code: %v resp: %v", resp.StatusCode, resp)
	}

	// read this byte by byte so you can show progress.
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(util.SegmentFilePath(d.SessionID, segmentID), bytes, os.ModePerm)
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

func (d *Dispatcher) acceptedStatusCodes(code int) bool {
	table := map[int]string{
		200: "OK",
		206: "Partial content",
	}
	return table[code] != ""
}
