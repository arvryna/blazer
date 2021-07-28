package data

import (
	"flag"
	"fmt"
)

const DEFAULT_THREAD_COUNT = 3
const DEFAULT_OUTPUT_PATH = "."

type CLIFlags struct {
	Url        string
	Thread     int
	OutputPath string
	Verbose    bool
}

func GetFilenameFromURL() {
}

func ParseCLIFlags() *CLIFlags {
	url := flag.String("url", "", "a string")
	out := flag.String("out", DEFAULT_OUTPUT_PATH, "a string")
	thread := flag.Int("thread", DEFAULT_THREAD_COUNT, "a number")
	// if *url == "" { // use regex and do proper analysis
	// 	fmt.Println("not valid URL")
	// 	return
	// }
	flag.Parse()

	cliFlags := CLIFlags{}
	cliFlags.Url = *url
	cliFlags.OutputPath = *out
	cliFlags.Thread = *thread
	return &cliFlags
}

func GetFormattedSize(size float64) string {
	mem := [6]string{"b", "kb", "mb", "gb", "tb", "pb"}
	i := 0
	for {
		if size < 1024 {
			return fmt.Sprintf("%.2f", size) + " " + mem[i]
		} else {
			size = size / 1024
			i++
		}
	}
}
