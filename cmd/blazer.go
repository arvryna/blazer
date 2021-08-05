package cmd

import (
	"fmt"
	"time"

	"github.com/arvyshka/blazer/internals"
	"github.com/arvyshka/blazer/internals/network"
	pkg "github.com/arvyshka/blazer/pkg/data"
)

func Start() {
	flags := CLIFlags{}
	err := flags.Parse()
	if err != nil {
		fmt.Println("Error parsing flags: ", err)
		return
	}
	if flags.Version {
		fmt.Println("Blazer version: ", internals.Version)
		return
	}
	setup(&flags)
}

// Life cycle of the app.
func setup(flags *CLIFlags) {
	fmt.Println("Fetching file meta..")
	meta := network.FileMeta{}
	err := meta.Fetch(flags.URL)
	if err != nil {
		fmt.Println("Can't initiate download", err)
		return
	}

	// Logging important info to user
	fmt.Println("File size: " + pkg.GetFormattedSize(meta.ContentLength))

	// Generate session ID for current download
	internals.SessionID = pkg.GenHash(flags.URL, flags.Thread)

	if pkg.FileExists(meta.FileName) {
		// FIX: Also check the File size, just to be sure that it wasn't an incomplete download.
		fmt.Println("File already exists, skipping download")
		return
	}

	// Using a temp folder in current dir to manage use artifacts of download.
	tempFileDir := internals.TempDirectory(internals.SessionID)
	if pkg.FileExists(tempFileDir) {
		fmt.Println("Resuming download..")
	} else {
		pkg.CreateDir(internals.TempDirectory(internals.SessionID), ".")
	}

	initiateDownload(flags, &meta)

	// Check file integrity
	if flags.Checksum != "" {
		res := pkg.FileIntegrityCheck("sha256", meta.FileName, flags.Checksum)
		fmt.Println("File integrity: ", res)
	}
	pkg.DeleteFile(internals.TempDirectory(internals.SessionID))
}

func initiateDownload(flags *CLIFlags, meta *network.FileMeta) {
	start := time.Now()

	path := flags.OutputPath
	if path == "" {
		path = meta.FileName
	}

	fmt.Println("Outputfile name: " + path)
	network.ConcurrentDownloader(meta, flags.Thread, path)
	fmt.Println("Download finished in: ", time.Since(start))
}
