package data

import (
	"crypto/rand"
	"flag"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"
)

type CLIFlags struct {
	Url        string
	Thread     int
	OutputPath string
	Verbose    bool
	Checksum   string
	Version    bool
}

func GenHash(s string) string {
	hash := fnv.New32a() // why not New64 ?
	hash.Write([]byte(s))
	return fmt.Sprintf("%v", hash.Sum32())
}

//generate random string
func GenRandomString(len int) string {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

// Delete file/folder
func DeleteFile(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDir(folderName string, dirPath string) {
	newpath := filepath.Join(".", folderName)
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating director")
	}
}

func ParseCLIFlags() *CLIFlags {
	ver := flag.Bool("v", false, "prints current version of blazer")
	url := flag.String("url", "", "Valid URL to download")
	out := flag.String("out", DEFAULT_OUTPUT_PATH, "output path to store the downloaded file")
	t := flag.Int("t", DEFAULT_THREAD_COUNT, "Thread count - Number of concurrent downloads")
	checksum := flag.String("checksum", "", "checksum SHA to verify file")
	// if *url == "" { // use regex and do proper analysis
	// 	fmt.Println("not valid URL")
	// 	return
	// }
	flag.Parse()

	cliFlags := CLIFlags{
		Url:        *url,
		OutputPath: *out,
		Thread:     *t,
		Checksum:   *checksum,
		Version:    *ver,
	}
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
