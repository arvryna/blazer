package data

import "fmt"

const TEMP_DIRECTORY = ".blazer_temp"
const DEFAULT_THREAD_COUNT = 10
const DEFAULT_OUTPUT_PATH = "."

const MEM_UNIT = 1024

func MemoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
}

func SegmentFilePath(id int) string {
	return fmt.Sprintf("%v/s-%v", TEMP_DIRECTORY, id)
}
