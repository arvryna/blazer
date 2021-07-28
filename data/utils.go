package data

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/arvpyrna/blazer/network"
)

const DEFAULT_THREAD_COUNT = 3
const DEFAULT_OUTPUT_PATH = "."

func GetFilenameFromURL() {
}

func ParseCLIFlags() {
	url := flag.String("url", "", "a string")
	outputPath := flag.String("out", DEFAULT_OUTPUT_PATH, "a string")
	thread := flag.Int("thread", DEFAULT_THREAD_COUNT, "a number")
	// if *url == "" { // use regex and do proper analysis
	// 	fmt.Println("not valid URL")
	// 	return
	// }
	flag.Parse()

	fmt.Println("Output path" + *outputPath)
	meta := network.GetFileMeta(*url)
	// fmt.Println("Content length of the file: " + strconv.Itoa(meta.ContentLength))
	fmt.Println("Download the file in " + strconv.Itoa(*thread) + " threads")
	fmt.Println("File size: " + getFormattedSize(meta.ContentLength))

	network.Download(*url, *thread)
}

func getFormattedSize(size float64) string {
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
