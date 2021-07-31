package data

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
	// finally check if the SHA matches and also check if the content length matches with
}
