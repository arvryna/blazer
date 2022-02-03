package cmd

import (
	"fmt"

	"github.com/arvryna/blazer/internals/cflags"
	"github.com/arvryna/blazer/internals/downloader"
)

/*
	This file handles unpacking essential configs, parsing flags
	and initiating the necessary download
*/

const (
	version = "0.4-beta"
)

var (
	flags cflags.CLIFlags
)

func Execute() {
	flags = cflags.CLIFlags{}

	err := flags.Parse()
	if err != nil {
		fmt.Println("Error parsing flags: ", err)
		return
	}

	flags.PerformEssentialChecks()
	downloader.Run(flags)
}
