package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"strings"
)

const MemUnit = 1024

func genChecksumSha256(path string) string {
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

func FileIntegrityCheck(hashFunc string, path string, expected string) bool {
	if strings.ToLower(hashFunc) == "sha256" {
		return (expected == genChecksumSha256(path))
	}
	fmt.Println(hashFunc, ": not implemented yet")
	return false
}

func GenHash(s string, threadCount int) string {
	hash := fnv.New32a() // why not New64 ?
	hash.Write([]byte(s))
	return fmt.Sprintf("%v-%v", hash.Sum32(), threadCount)
}

func MemoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
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
