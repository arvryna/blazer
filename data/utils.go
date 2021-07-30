package data

import (
	"flag"
	"fmt"
)

type CLIFlags struct {
	Url        string
	Thread     int
	OutputPath string
	Verbose    bool
	Checksum   string
}

func GetFilenameFromURL() {
}

func ParseCLIFlags() *CLIFlags {
	url := flag.String("url", "", "a string")
	out := flag.String("out", DEFAULT_OUTPUT_PATH, "a string")
	thread := flag.Int("thread", DEFAULT_THREAD_COUNT, "a number")
	checksum := flag.String("checksum", "", "checksum SHA to verify file")
	// if *url == "" { // use regex and do proper analysis
	// 	fmt.Println("not valid URL")
	// 	return
	// }
	flag.Parse()

	cliFlags := CLIFlags{}
	cliFlags.Url = *url
	cliFlags.OutputPath = *out
	cliFlags.Thread = *thread
	cliFlags.Checksum = *checksum
	return &cliFlags
}

func GetFormattedSize(size float64) string {
	mem := MemoryFormatStrings()
	i := 0
	for {
		if size < MEM_UNIT {
			return fmt.Sprintf("%.2f", size) + " " + mem[i]
		} else {
			size = size / MEM_UNIT
			i++
		}
	}
}
