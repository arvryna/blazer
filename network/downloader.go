package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

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
	request, _ := BuildRequest(http.MethodGet, meta.FileUrl)
	chunks := data.CalculateChunks(int(meta.ContentLength), thread)
	var wg sync.WaitGroup
	for i, segment := range chunks.Segments {
		// start before concurrency
		wg.Add(1)
		// capturing values as they change
		i := i
		segment := segment
		request := request
		go func() {
			defer wg.Done() // defer is good pattern than trying to close in a specific place
			DownloadSegment(request, i, segment)
		}()
	}
	wg.Wait()
	mergefiles(chunks)
	// data, _ := ioutil.ReadAll(resp.Body)
	// ioutil.WriteFile("out/test.pdf", data, os.ModePerm)
}

func DownloadSegment(request *http.Request, i int, r data.Range) {
	fmt.Printf("\nstarting segment : " + fmt.Sprintf("# %v [%v-%v]", i, r.Start, r.End))
	request.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", r.Start, r.End))
	resp, _ := http.DefaultClient.Do(request)
	// handle error
	// read this byte by byte so you can show progress
	data, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(fmt.Sprintf("out/segment-%v.pdf", i), data, os.ModePerm)
	fmt.Println("Downloaded segment: ", i)
}

func mergefiles(chunks *data.Chunks) {
	f, _ := os.OpenFile("out/fin.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer f.Close()
	for i := range chunks.Segments {
		data, err := ioutil.ReadFile(fmt.Sprintf("out/segment-%v.pdf", i))
		if err != nil {
			fmt.Println(err)
		}
		n, err := f.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nbytes merged: %v of segment: %v", n, i)
	}
}
