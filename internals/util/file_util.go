package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	MemUnit         = 1024
	SHA256Algorithm = "sha256"
	MD5Algorithm    = "md5"
)

// Create a directory with session ID, Session ID is hash of URL and threadcount.
func TempDirectory(session string) string {
	return fmt.Sprintf(".blazer_temp-%v", session)
}

/*
* Segments are stored in side the temproary directory above,
* there are n segments, n represents threadcount. if thread = 10
* there will be 10 segments in the temp folder.
 */
func SegmentFilePath(session string, fileID int) string {
	return fmt.Sprintf("%v/s-%v", TempDirectory(session), fileID)
}

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func DeleteFile(name string) {
	err := os.RemoveAll(name)
	if err != nil {
		log.Fatal(err)
	}
}

// Create dir in current directory.
func CreateDir(folderName string, dirPath string) {
	newpath := filepath.Join(".", folderName)
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory")
	}
}

func FileIntegrityCheck(hashFunc string, path string, expected string) bool {
	if strings.ToLower(hashFunc) == SHA256Algorithm {
		return (expected == genChecksumSha256(path))
	}
	fmt.Println(hashFunc, ": not implemented yet")
	return false
}

func GenHash(s string, threadCount int) string {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return fmt.Sprintf("%v-%v", hash.Sum32(), threadCount)
}

func GetFormattedSize(size float64) string {
	i := 0
	mem := memoryFormatStrings()
	for {
		if size < MemUnit {
			return fmt.Sprintf("%.2f", size) + " " + mem[i]
		}
		size /= MemUnit
		i++
	}
}

// **** Private Functions ****

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

func memoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
}
