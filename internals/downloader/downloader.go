package downloader

import (
	"fmt"
	"time"

	"github.com/arvryna/blazer/internals/cflags"
	"github.com/arvryna/blazer/internals/network"
	"github.com/arvryna/blazer/internals/util"
)

const optimized_Download_Unsupported = "Optimized downloading not supported by server!"

type Downloader struct {
	flags cflags.CLIFlags
}

func Run(flags cflags.CLIFlags) {
	fmt.Println("Fetching file meta..")
	meta := network.FileMeta{}
	err := meta.Fetch(flags.URL)
	if err != nil {
		fmt.Println("Can't initiate download", err)
		return
	}

	if doesServerSupportRangeHeader(&meta) {
		flags.Thread = 1
		fmt.Println(optimized_Download_Unsupported)
	}

	// Generate session ID for current download
	sessionID := util.GenHash(flags.URL, flags.Thread)
	manageDownloadFlow(&flags, &meta, sessionID)
}

func manageDownloadFlow(flags *cflags.CLIFlags, meta *network.FileMeta, sessionID string) {

	// Logging important info to user
	fmt.Println("File size: " + util.GetFormattedSize(meta.ContentLength))

	if util.FileExists(meta.FileName) {
		// FIX: Also check the File size, just to be sure that it wasn't an incomplete download.
		fmt.Println("File already exists, skipping download")
		return
	}

	// Using a temp folder in current dir to manage use artifacts of download.
	tempFileDir := util.TempDirectory(sessionID)
	if util.FileExists(tempFileDir) {
		fmt.Println("Resuming download..")
	} else {
		util.CreateDir(util.TempDirectory(sessionID), ".")
	}

	isDownloadComplete := downloadAndMerge(flags, meta, sessionID)

	if isDownloadComplete {
		// Perform file integrity check
		if flags.Checksum != "" {
			res := util.FileIntegrityCheck("sha256", meta.FileName, flags.Checksum)
			fmt.Println("File integrity: ", res)
		}
		util.DeleteFile(util.TempDirectory(sessionID))
	}
}

func downloadAndMerge(flags *cflags.CLIFlags, meta *network.FileMeta, sessionID string) bool {
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
