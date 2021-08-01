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
	meta, err := network.GetFileMeta(flags.Url)
	if err != nil {
		fmt.Println("Can't initiate download", err)
		return
	}

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
		fmt.Println("File integrity: ", res)
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
	fmt.Println("Download finished in: ", time.Since(start))
}

func main() {
	flags, err := data.ParseCLIFlags()
	if err != nil {
		fmt.Println("Error parsing flags: ", err)
		return
	}
	if flags.Version {
		fmt.Println("Blazer version: ", data.VERSION)
		return
	}
	setup(flags)
}
