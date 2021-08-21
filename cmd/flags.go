package cmd

import (
	"errors"
	"flag"

	"github.com/arvyshka/blazer/internals/network"
)

type CLIFlags struct {
	URL        string
	Thread     int
	OutputPath string
	Verbose    bool
	Checksum   string
	Version    bool
}

// Default number of threads used for download if user don't specify thread count.
const DefaultThreadCount = 10

func (f *CLIFlags) Parse() error {
	ver := flag.Bool("v", false, "Prints current version of blazer")
	urlString := flag.String("url", "", "Valid URL to download")
	out := flag.String("out", "", "Output path to store the downloaded file")
	t := flag.Int("t", DefaultThreadCount, "Thread count - Number of concurrent downloads")
	checksum := flag.String("checksum", "", "Checksum SHA256(currently supported) to verify file")

	flag.Parse()

	if *urlString == "" {
		return errors.New("url is mandatory")
	}

	if !network.IsValidURL(*urlString) {
		return errors.New("invalid URL")
	}

	f.Version = *ver
	f.URL = *urlString
	f.OutputPath = *out
	f.Thread = *t
	f.Checksum = *checksum

	return nil
}
