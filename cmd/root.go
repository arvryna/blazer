package cmd

import (
	"fmt"

	"github.com/arvryna/blazer/internals/cflags"
	"github.com/arvryna/blazer/internals/downloader"
)

// This file parses the CLI flags and initiates download and initiating the necessary download

func Execute() {
	flags := cflags.CLIFlags{}

	err := flags.Parse()
	if err != nil {
		fmt.Println("Error parsing flags: ", err)
		return
	}

	flags.PerformEssentialChecks()

	d := downloader.Downloader{
		Flags: flags,
	}

	d.Run(flags)
}
