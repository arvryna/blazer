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
	fmt.Println("File size: " + data.GetFormattedSize(meta.ContentLength))

	network.Download(flags.Url, flags.Thread)
}

func main() {
	r := data.CalculateChunks(100, 7)
	fmt.Println(r)
	setup()
}
