package main

import (
	"fmt"

	"github.com/arvryna/blazer/internal/cflags"
	"github.com/arvryna/blazer/internal/downloader"
	"github.com/arvryna/blazer/internal/util"
)

func execute() {
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

func main() {
	execute()
}
