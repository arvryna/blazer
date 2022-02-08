package cflags

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/arvryna/blazer/internals/network"
)

type CLIFlags struct {
	URL        string
	Thread     int
	OutputPath string
	Verbose    bool
	Checksum   string
	Version    bool
}

const (
	DefaultThreadCount = 10
	version            = "0.5-beta"
)

func (f *CLIFlags) Parse() error {
	ver := flag.Bool("v", false, "Prints current version of blazer")
	urlString := flag.String("url", "", "Valid URL to download")
	out := flag.String("out", "", "Output path to store the downloaded file")
	t := flag.Int("t", DefaultThreadCount, "Thread count - Number of concurrent downloads")
	checksum := flag.String("checksum", "", "Checksum SHA256(currently supported) to verify file")
	flag.Parse()

	f.Version = *ver
	f.URL = *urlString
	f.OutputPath = *out
	f.Thread = *t
	f.Checksum = *checksum

	return nil
}

func (f *CLIFlags) HasValidDownloadURL() (bool, error) {
	if !network.IsValidURL(f.URL) {
		return false, errors.New("Invalid URL, a valid URL is mandatory, pass URL using -url flag")
	}
	return true, nil
}

func (f *CLIFlags) PerformEssentialChecks() {
	if f.Version {
		fmt.Println("Blazer version: ", version)
		os.Exit(0)
	}

	ok, err := f.HasValidDownloadURL()
	if !ok {
		fmt.Println(err)
		os.Exit(1)
	}
}
