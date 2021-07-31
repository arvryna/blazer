package main

import (
	"fmt"
	"time"

	"github.com/arvpyrna/blazer/data"
	"github.com/arvpyrna/blazer/network"
)

// Life cycle of the app
func setup(flags *data.CLIFlags) {
	meta := network.GetFileMeta(flags.Url)

	// Logging important info to user
	logInfo(flags, meta)

	// Generate session ID for current download
	data.SESSION_ID = data.GenHash(flags.Url, flags.Thread)

	// Using a temp folder in current dir to manage use artifacts of download
	data.CreateDir(data.TempDirectory(data.SESSION_ID), ".")
	initiateDownload(flags, meta)
	data.DeleteFile(data.TempDirectory(data.SESSION_ID))
}

func logInfo(flags *data.CLIFlags, meta *network.FileMeta) {
	fmt.Println("Fetching file meta..")
	fmt.Printf("Download the file in %v threads", flags.Thread)
	fmt.Println("\nFile size: " + data.GetFormattedSize(meta.ContentLength))
}

func initiateDownload(flags *data.CLIFlags, meta *network.FileMeta) {
	start := time.Now()
	network.ConcurrentDownloader(meta, flags.Thread)
	elapsed := time.Since(start)
	fmt.Printf("Download finished in: %v", elapsed)
}

func main() {
	flags := data.ParseCLIFlags()
	if flags.Version {
		fmt.Printf("Blazer version: [%v]\n", data.VERSION)
		return
	}
	setup(flags)
}
