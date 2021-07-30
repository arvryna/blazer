package main

import (
	"fmt"
	"time"

	"github.com/arvpyrna/blazer/data"
	"github.com/arvpyrna/blazer/network"
)

func setup() {
	// create folder if not exists
	// delete that folder after download is successful
	flags := data.ParseCLIFlags()

	fmt.Println("Fetching file meta..")
	meta := network.GetFileMeta(flags.Url)

	fmt.Printf("Download the file in %v threads", flags.Thread)
	fmt.Println("\nFile size: " + data.GetFormattedSize(meta.ContentLength))

	start := time.Now()
	network.ConcurrentDownloader(meta, flags.Thread)
	elapsed := time.Since(start)
	fmt.Printf("Download finished in: %v", elapsed)
}

func main() {
	data.CreateDir(data.TEMP_DIRECTORY, ".")
	setup()
	data.DeleteFile(data.TEMP_DIRECTORY)
}
