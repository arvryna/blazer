package data

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"strings"
)

type CLIFlags struct {
	Url        string
	Thread     int
	OutputPath string
	Verbose    bool
	Checksum   string
	Version    bool
}

func FileIntegrityCheck(hashFunc string, path string, expected string) bool {
	if strings.ToLower(hashFunc) == "sha256" {
		return (expected == GenChecksumSha256(path))
	} else {
		fmt.Println(hashFunc, ": not implemented yet")
		return false
	}
}

func GenChecksumSha256(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		fmt.Println(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func GenHash(s string, threadCount int) string {
	hash := fnv.New32a() // why not New64 ?
	hash.Write([]byte(s))
	return fmt.Sprintf("%v-%v", hash.Sum32(), threadCount)
}

func ParseCLIFlags() *CLIFlags {
	ver := flag.Bool("v", false, "prints current version of blazer")
	url := flag.String("url", "", "Valid URL to download")
	out := flag.String("out", "", "output path to store the downloaded file")
	t := flag.Int("t", DEFAULT_THREAD_COUNT, "Thread count - Number of concurrent downloads")
	checksum := flag.String("checksum", "", "checksum SHA256(currently supported) to verify file")
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
