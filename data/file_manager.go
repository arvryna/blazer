package data

import (
	"fmt"
	"io/ioutil"
	"os"
)

func MergeFiles(chunks *Chunks) {
	f, _ := os.OpenFile("out/fin.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer f.Close()
	for i := range chunks.Segments {
		fileName := fmt.Sprintf("out/s-%v.pdf", i)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write(data)
		if err != nil {
			fmt.Println(err)
		}
		err = os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
		}
	}
	// finally check if the SHA matches and also check if the content length matches with
	// bytes merged
}
