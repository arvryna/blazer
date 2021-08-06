package internals

// File for storing global constants and functions

import "fmt"

var SessionID string

// Constant to track the current version of CLI.
const Version = "0.3-beta"

// Default number of threads used for download if user don't specify thread count.
const DefaultThreadCount = 10

const MemUnit = 1024

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

func MemoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
}
