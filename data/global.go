package data

import "fmt"

var SessionID string

// Constants.
const Version = "0.3-beta"
const DefaultThreadCount = 10
const MemUnit = 1024

// Functions.
func TempDirectory(session string) string {
	return fmt.Sprintf(".blazer_temp-%v", session)
}

func SegmentFilePath(session string, fileID int) string {
	return fmt.Sprintf("%v/s-%v", TempDirectory(session), fileID)
}

func MemoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
}
