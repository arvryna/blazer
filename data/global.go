package data

import "fmt"

var SESSION_ID = GenRandomString(7)

// Constants
const VERSION = "0.1-alpha"
const TEMP_DIRECTORY = ".blazer_temp"
const DEFAULT_THREAD_COUNT = 10
const DEFAULT_OUTPUT_PATH = "."
const MEM_UNIT = 1024

// Functions
func MemoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
}

func SegmentFilePath(id int) string {
	return fmt.Sprintf("%v/s-%v", TEMP_DIRECTORY, id)
}
