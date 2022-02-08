package cmd

import (
	"fmt"

	"github.com/arvryna/blazer/internals/cflags"
	"github.com/arvryna/blazer/internals/downloader"
	"github.com/arvryna/blazer/internals/util"
)

// This file parses the CLI flags and initiates download

func Execute() {
	flags := cflags.CLIFlags{}

	err := flags.Parse()
	if err != nil {
		fmt.Println("Error parsing flags: ", err)
		return
	}

	flags.PerformEssentialChecks()

	sessionID := util.GenHash(flags.URL, flags.Thread)

	d := downloader.Downloader{
		Flags:     flags,
		SessionID: sessionID,
	}

	d.Run()
}
