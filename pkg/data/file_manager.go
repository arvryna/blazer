package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// Delete file/folder.
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
