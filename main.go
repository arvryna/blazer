package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/arvpyrna/blazer/data"
	"github.com/arvpyrna/blazer/network"
)

func setup() {
	flags := data.ParseCLIFlags()

	fmt.Println("Fetching file meta..")
	meta := network.GetFileMeta(flags.Url)
	fmt.Println("Output path" + flags.OutputPath)
	fmt.Println("Download the file in " + strconv.Itoa(flags.Thread) + " threads")
	fmt.Println("content-length: " + fmt.Sprintf("%f", meta.ContentLength))
	fmt.Println("File size: " + data.GetFormattedSize(meta.ContentLength))

	start := time.Now()
	network.ConcurrentDownloader(meta, flags.Thread)
	elapsed := time.Since(start)
	log.Println(fmt.Sprintf("Download finished in: %s", elapsed))
}

func main() {
	setup()
}
