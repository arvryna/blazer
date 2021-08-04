package data

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"hash/fnv"
	"io"
	"net/url"
	"os"
	"strings"
)

type CLIFlags struct {
	URL        string
	Thread     int
	OutputPath string
	Verbose    bool
	Checksum   string
	Version    bool
}

func FileIntegrityCheck(hashFunc string, path string, expected string) bool {
	if strings.ToLower(hashFunc) == "sha256" {
		return (expected == GenChecksumSha256(path))
	}
	fmt.Println(hashFunc, ": not implemented yet")
	return false
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

func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func ParseCLIFlags() (*CLIFlags, error) {
	ver := flag.Bool("v", false, "Prints current version of blazer")
	urlString := flag.String("url", "", "Valid URL to download")
	out := flag.String("out", "", "Output path to store the downloaded file")
	t := flag.Int("t", DefaultThreadCount, "Thread count - Number of concurrent downloads")
	checksum := flag.String("checksum", "", "Checksum SHA256(currently supported) to verify file")
	flag.Parse()

	if *urlString == "" {
		return nil, errors.New("url is mandatory")
	}

	if !IsValidURL(*urlString) {
		return nil, errors.New("invalid URL")
	}

	cliFlags := CLIFlags{
		URL:        *urlString,
		OutputPath: *out,
		Thread:     *t,
		Checksum:   *checksum,
		Version:    *ver,
	}
	return &cliFlags, nil
}

func GetFormattedSize(size float64) string {
	mem := MemoryFormatStrings()
	i := 0
	for {
		if size < MemUnit {
			return fmt.Sprintf("%.2f", size) + " " + mem[i]
		}
		size = size / MemUnit
		i++
	}
}
