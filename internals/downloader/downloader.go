package downloader

import (
	"fmt"
	"os"
	"time"

	"github.com/arvryna/blazer/internals/cflags"
	"github.com/arvryna/blazer/internals/network"
	"github.com/arvryna/blazer/internals/util"
)

const optimizedDownloadUnsupported = "Optimized downloading not supported by server!, File will be downloaded with a single thread"

type Downloader struct {
	SessionID string
	Flags     cflags.CLIFlags
}

func (d *Downloader) Run() {
	fmt.Println("Fetching file meta..")

	meta := network.FileMeta{}
	err := meta.Fetch(d.Flags.URL)
	if err != nil {
		fmt.Println("Can't initiate download, file meta info can't be fetched", err)
		return
	}

	d.performEssentialChecks(&meta)
	d.createTempDirectory()

	fmt.Println("File size: " + util.GetFormattedSize(meta.ContentLength))

	// Using a temp folder in current dir to manage use artifacts of download.
	isDownloadComplete := d.downloadAndMerge(&meta)

	if isDownloadComplete {
		if d.Flags.Checksum != "" {
			res := util.FileIntegrityCheck("sha256", meta.FileName, d.Flags.Checksum)
			fmt.Println("File integrity: ", res)
		}
		util.DeleteFile(util.TempDirectory(d.SessionID))
	}
}

func (d *Downloader) performEssentialChecks(meta *network.FileMeta) {

	if util.FileExists(meta.FileName) {
		// TODO: Also check the File size, just to be sure that it wasn't an incomplete download.
		fmt.Println("File already exists, skipping download")
		os.Exit(1)
	}

	// Not all servers can allow multi-threaded file download, fallback to single threaded download when necessary
	if d.doesServerSupportRangeHeader(meta) {
		d.Flags.Thread = 1
		fmt.Println(optimizedDownloadUnsupported)
	}
}

func (d *Downloader) createTempDirectory() {
	tempFileDir := util.TempDirectory(d.SessionID)
	if util.FileExists(tempFileDir) {
		fmt.Println("Resuming download..")
	} else {
		util.CreateDir(util.TempDirectory(d.SessionID), ".")
	}
}

func (d *Downloader) downloadAndMerge(meta *network.FileMeta) bool {
	fmt.Println("Download the file in threads: ", d.Flags.Thread)
	outputPath := d.Flags.OutputPath
	if outputPath == "" {
		outputPath = meta.FileName
	}

	fmt.Println("Outputfile name: " + outputPath)

	start := time.Now()

	dispatcher := network.Dispatcher{
		Meta:        meta,
		ThreadCount: d.Flags.Thread,
		SessionID:   d.SessionID,
	}

	chunks, isDownloadComplete := dispatcher.InitiateConcurrentDispatch()

	if isDownloadComplete {
		fmt.Println("Download finished in: ", time.Since(start))
		fmt.Println("Merging downloaded files...")
		err := chunks.Merge(outputPath, d.SessionID)
		if err != nil {
			fmt.Println("File merging failed ", err)
		}
	} else {
		fmt.Println("Download failed: Some segments were not downloaded, please re-intiate the download")
	}

	return isDownloadComplete
}

func (d *Downloader) doesServerSupportRangeHeader(meta *network.FileMeta) bool {
	return meta.AcceptRanges != "bytes"
}
