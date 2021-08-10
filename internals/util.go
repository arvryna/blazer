package internals

// File for storing global constants and functions

import "fmt"

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
