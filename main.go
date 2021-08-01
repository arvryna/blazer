package main

import (
	"fmt"
	"time"

	"github.com/arvpyrna/blazer/data"
	"github.com/arvpyrna/blazer/network"
)

// Life cycle of the app
func setup(flags *data.CLIFlags) {
	fmt.Println("Fetching file meta..")
	meta := network.GetFileMeta(flags.Url)

	// Logging important info to user
	fmt.Println("File size: " + data.GetFormattedSize(meta.ContentLength))

	// Generate session ID for current download
	data.SESSION_ID = data.GenHash(flags.Url, flags.Thread)

	// Using a temp folder in current dir to manage use artifacts of download
	tempFileDir := data.TempDirectory(data.SESSION_ID)
	if data.FileExists(tempFileDir) {
		fmt.Println("Resuming download..")
	} else {
		data.CreateDir(data.TempDirectory(data.SESSION_ID), ".")
	}
	initiateDownload(flags, meta)

	// Check file integrity
	if flags.Checksum != "" {
		res := data.FileIntegrityCheck("sha256", meta.FileName, flags.Checksum)
		fmt.Printf("\nFile integrity: %v\n", res)
	}
	data.DeleteFile(data.TempDirectory(data.SESSION_ID))
}

func initiateDownload(flags *data.CLIFlags, meta *network.FileMeta) {
	start := time.Now()

	path := flags.OutputPath
	if path == "" {
		path = meta.FileName
	}

	fmt.Println("Outputfile name: " + path)

	network.ConcurrentDownloader(meta, flags.Thread, path)
	fmt.Printf("Download finished in: %v", time.Since(start))
}

func main() {
	flags := data.ParseCLIFlags()
	if flags.Version {
		fmt.Printf("Blazer version: [%v]\n", data.VERSION)
		return
	}
	setup(flags)
}
