package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/arvyshka/blazer/data"
)

func ConcurrentDownloader(meta *FileMeta, thread int, outputName string) {
	fmt.Println("Download the file in threads: ", thread)
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
			err = DownloadSegment(request, i, segment)
			if err != nil {
				fmt.Println("Download segment failed: ", i, err)
			}
		}()
	}
	wg.Wait()

	// Do not merge file, if download has filed
	// if the file to download is already there, you can skip the download
	err := data.MergeFiles(chunks, outputName)
	if err != nil {
		fmt.Println("File merging failed ", err)
	}
}

func acceptedStatusCodes(code int) bool {
	table := map[int]string{
		200: "OK",
		206: "Partial content",
	}
	return table[code] != ""
}

func DownloadSegment(request *http.Request, i int, r data.Range) error {
	request.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r.Start, r.End))
	resp, err := HTTPClient().Do(request)

	if !acceptedStatusCodes(resp.StatusCode) {
		return fmt.Errorf("received [un-expected] status code: %v resp: %v", resp.StatusCode, resp)
	}

	if err != nil {
		return err
	}

	// read this byte by byte so you can show progress
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(data.SegmentFilePath(data.SESSION_ID, i), bytes, os.ModePerm)
	if err != nil {
		return err
	}

	// Verify if segment is downloaded successfully
	if len(bytes) == int(resp.ContentLength) {
		fmt.Println("Downloaded segment: ", i)
	} else {
		return fmt.Errorf("incomplete segment: %v content-len %v", i, resp.ContentLength)
	}
	return nil
}
