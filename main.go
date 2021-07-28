package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

// storing data in current drive
// show progress
// use interactive terminal to show progress across each thread, so it more cooler
// upload it to free file sharing service upto some MB, you can build one for yourself and share
// it with others, write in go, or node js and also use it here, a simple service
// the life of the service can be just 1 hour, there you can learn so many interesting things
// download files after giving username and password(with authentication)
// export as package too, later somehow
// show current download speed and ETA

const DEFAULT_THREAD_COUNT = 3
const DEFAULT_OUTPUT_PATH = "."

type FileMeta struct {
	ContentLength float64
	ServerName    string
	Age           int
	ContentType   string
}

func getFormattedSize(size float64) string {
	mem := [6]string{"b", "kb", "mb", "gb", "tb", "pb"}
	i := 0
	for {
		if size < 1024 {
			return fmt.Sprintf("%.2f", size) + " " + mem[i]
		} else {
			size = size / 1024
			i++
		}
	}
}

func preferredChunks(contentLength int) {

}

func getFileMeta(url string) *FileMeta {
	//find content length
	r, err := http.NewRequest("HEAD", url, nil)
	r.Header.Set("User-Agent", "Blazer")
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
	return &meta
}

func setup() {
	url := flag.String("url", "", "a string")
	outputPath := flag.String("out", DEFAULT_OUTPUT_PATH, "a string")
	thread := flag.Int("thread", DEFAULT_THREAD_COUNT, "a number")
	// if *url == "" { // use regex and do proper analysis
	// 	fmt.Println("not valid URL")
	// 	return
	// }
	flag.Parse()
	fmt.Println("Going to download resource: " + *url + " Thread: " + strconv.Itoa(*thread))
	fmt.Println("Output path" + *outputPath)

	meta := getFileMeta(*url)
	// fmt.Println("Content length of the file: " + strconv.Itoa(meta.ContentLength))
	fmt.Println("file meta details: ", meta)
	fmt.Println("Download the file in" + strconv.Itoa(*thread) + "parts")
	fmt.Println("File size: " + getFormattedSize(meta.ContentLength))
}

func main() {
	setup()
}
