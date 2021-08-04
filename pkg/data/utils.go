package data

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"io"
	"net/url"
	"os"
	"strings"
)

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

func GetFormattedSize(size float64) string {
	i := 0
	mem := MemoryFormatStrings()
	for {
		if size < MemUnit {
			return fmt.Sprintf("%.2f", size) + " " + mem[i]
		}
		size /= MemUnit
		i++
	}
}
