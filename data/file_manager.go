package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
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
		fmt.Println("Error creating directory")
	}
}

func MergeFiles(chunks *Chunks, outputName string) {
	println("Merging files..")
	f, err := os.OpenFile(outputName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer f.Close()

	bytesMerged := 0
	for i := range chunks.Segments {
		fileName := SegmentFilePath(SESSION_ID, i)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
		}
		bytes, err := f.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		err = os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
		}
		bytesMerged += bytes
	}

	// Check if download complete
	if bytesMerged == chunks.TotalSize {
		fmt.Println("File downloaded successfully..")
	} else {
		fmt.Println("File download is incomplete, retry")
	}
	// check if the SHA matches and also check if the content length matches with
}
