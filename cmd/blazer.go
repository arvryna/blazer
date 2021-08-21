package cmd

import (
	"fmt"
	"time"

	"github.com/arvyshka/blazer/internals"
	"github.com/arvyshka/blazer/internals/network"
	pkg "github.com/arvyshka/blazer/pkg/data"
)

// Constant to track the current version of CLI.
const Version = "0.4-beta"
const (
	Optimized_Download_Unsupported = "Optimized downloading not supported by server!"
)

func Start() {
	flags := CLIFlags{}
	err := flags.Parse()
	if err != nil {
		fmt.Println("Error parsing flags: ", err)
		return
	}

	// If the user only want to check version, show version info and exit
	if flags.Version {
		fmt.Println("Blazer version: ", Version)
		return
	}

	fmt.Println("Fetching file meta..")
	meta := network.FileMeta{}
	err = meta.Fetch(flags.URL)
	if err != nil {
		fmt.Println("Can't initiate download", err)
		return
	}

	if doesServerSupportRangeHeader(&meta) {
		flags.Thread = 1
		fmt.Println(Optimized_Download_Unsupported)
	}

	// Generate session ID for current download
	sessionID := pkg.GenHash(flags.URL, flags.Thread)
	manageDownloadFlow(&flags, &meta, sessionID)
}

func manageDownloadFlow(flags *CLIFlags, meta *network.FileMeta, sessionID string) {

	// Logging important info to user
	fmt.Println("File size: " + pkg.GetFormattedSize(meta.ContentLength))

	if pkg.FileExists(meta.FileName) {
		// FIX: Also check the File size, just to be sure that it wasn't an incomplete download.
		fmt.Println("File already exists, skipping download")
		return
	}

	// Using a temp folder in current dir to manage use artifacts of download.
	tempFileDir := internals.TempDirectory(sessionID)
	if pkg.FileExists(tempFileDir) {
		fmt.Println("Resuming download..")
	} else {
		pkg.CreateDir(internals.TempDirectory(sessionID), ".")
	}

	isDownloadComplete := downloadAndMerge(flags, meta, sessionID)

	if isDownloadComplete {
		// Perform file integrity check
		if flags.Checksum != "" {
			res := pkg.FileIntegrityCheck("sha256", meta.FileName, flags.Checksum)
			fmt.Println("File integrity: ", res)
		}
		pkg.DeleteFile(internals.TempDirectory(sessionID))
	}
}

func downloadAndMerge(flags *CLIFlags, meta *network.FileMeta, sessionID string) bool {
	fmt.Println("Download the file in threads: ", flags.Thread)
	outputPath := flags.OutputPath
	if outputPath == "" {
		outputPath = meta.FileName
	}

	fmt.Println("Outputfile name: " + outputPath)

	start := time.Now()
	chunks, isDownloadComplete := network.ConcurrentDownloader(meta, flags.Thread, sessionID)

	if isDownloadComplete {
		fmt.Println("Download finished in: ", time.Since(start))
		fmt.Println("Merging downloaded files...")
		err := chunks.Merge(outputPath, sessionID)
		if err != nil {
			fmt.Println("File merging failed ", err)
		}
	} else {
		fmt.Println("Download failed: Some segments were not downloaded, please re-intiate the download")
	}

	return isDownloadComplete
}

func doesServerSupportRangeHeader(meta *network.FileMeta) bool {
	return meta.AcceptRanges != "bytes"
}
