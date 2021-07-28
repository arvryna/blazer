package main

import (
	"fmt"
	"strconv"

	"github.com/arvpyrna/blazer/data"
	"github.com/arvpyrna/blazer/network"
)

func setup() {
	flags := data.ParseCLIFlags()

	meta := network.GetFileMeta(flags.Url)
	fmt.Println("Output path" + flags.OutputPath)
	fmt.Println("Download the file in " + strconv.Itoa(flags.Thread) + " threads")
	fmt.Println("content-length: " + fmt.Sprintf("%f", meta.ContentLength))
	fmt.Println("File size: " + data.GetFormattedSize(meta.ContentLength))

	network.ConcurrentDownloader(meta, flags.Thread)
}

func main() {
	setup()
}
