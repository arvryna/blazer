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

func MergeFiles(chunks *Chunks, outputName string) error {
	fmt.Println("Merging files..")
	f, err := os.OpenFile(outputName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	bytesMerged := 0
	for i := range chunks.Segments {
		fileName := SegmentFilePath(SESSION_ID, i)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			return err
		}
		bytes, err := f.Write(data)
		if err != nil {
			return err
		}
		err = os.Remove(fileName)
		if err != nil {
			return err
		}
		bytesMerged += bytes
	}

	if bytesMerged == chunks.TotalSize {
		fmt.Println("File downloaded successfully..")
	} else {
		return fmt.Errorf("File download is incomplete, retry")
	}
	return nil
}
