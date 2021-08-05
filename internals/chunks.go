package internals

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Range struct {
	Start int
	End   int
}

type Chunks struct {
	Size      int     // size of a single chunk
	TotalSize int     // size of the overall file
	Segments  []Range // segments [[0,n1],[n1+1,n1+chunkSize]....[,n]]
	Count     int     // number of chunks
}

// Compute the chunks for a given parts(thread count).
func (c *Chunks) ComputeChunks() {
	c.Size = int(float64(c.TotalSize) / float64(c.Count))
	pos := -1
	for i := 0; i < c.Count; i++ {
		r := Range{}
		r.Start = pos + 1
		pos += c.Size

		// Case 1
		if pos > c.TotalSize {
			// we have already divided enough segments, so can exit early.
			r.End = c.TotalSize
			c.Count = i + 1
			c.Segments = append(c.Segments, r)
			break
		}

		// Case 2
		if (i == c.Count-1) && pos < c.TotalSize {
			r.End = c.TotalSize
			c.Segments = append(c.Segments, r)
			break
		}
		r.End = pos
		c.Segments = append(c.Segments, r)
	}
}

// Merge all segments into a single file.
func (c *Chunks) Merge(outputName string) error {
	fmt.Println("Merging files..")
	f, err := os.OpenFile(outputName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	bytesMerged := 0
	for i := range c.Segments {
		fileName := SegmentFilePath(SessionID, i)
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

	if bytesMerged == c.TotalSize {
		fmt.Println("File downloaded successfully..")
	} else {
		return errors.New("file download is incomplete, retry")
	}
	return nil
}
