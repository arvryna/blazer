package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/arvyshka/blazer/internals"
	pkg "github.com/arvyshka/blazer/pkg/data"
)

func acceptedStatusCodes(code int) bool {
	table := map[int]string{
		200: "OK",
		206: "Partial content",
	}
	return table[code] != ""
}

// ConcurrentDownloader: Concurrently download the resource with specified concurrency value.
func ConcurrentDownloader(meta *FileMeta, thread int, outputName string) {
	fmt.Println("Download the file in threads: ", thread)
	chunks := internals.Chunks{Count: thread, TotalSize: int(meta.ContentLength)}
	chunks.ComputeChunks()
	var wg sync.WaitGroup
	for i, segment := range chunks.Segments {
		// if segment exist skip current segment download.
		if pkg.FileExists(internals.SegmentFilePath(internals.SessionID, i)) {
			// fmt.Println("Segment Id: ", i, "already downloaded").
			continue
		}
		request, err := BuildRequest(http.MethodGet, meta.FileURL)
		if err != nil {
			fmt.Println("Could not build request: ", err)
		}
		wg.Add(1)
		i := i
		segment := segment
		go func() {
			defer wg.Done()
			err = DownloadSegment(request, i, segment)
			if err != nil {
				fmt.Println("Download segment failed: ", i, err)
			}
		}()
	}
	wg.Wait()

	// Do not merge file, if download has filed
	err := chunks.Merge(outputName)
	if err != nil {
		fmt.Println("File merging failed ", err)
	}
}

// DownloadSegment: download a specific piece of the bytes of the file that we want to download.
func DownloadSegment(request *http.Request, segmentID int, r internals.Range) error {
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

	err = ioutil.WriteFile(internals.SegmentFilePath(internals.SessionID, segmentID), bytes, os.ModePerm)
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
