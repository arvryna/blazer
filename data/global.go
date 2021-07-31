package data

import "fmt"

var SESSION_ID string = ""

// Constants
const VERSION = "0.1-alpha"
const DEFAULT_THREAD_COUNT = 10
const DEFAULT_OUTPUT_PATH = "."
const MEM_UNIT = 1024

// Functions
func TempDirectory(session string) string {
	return fmt.Sprintf(".blazer_temp-%v", session)
}

func SegmentFilePath(session string, fileId int) string {
	return fmt.Sprintf("%v/s-%v", TempDirectory(session), fileId)
}

func MemoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
}
