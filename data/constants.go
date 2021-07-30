package data

const DEFAULT_THREAD_COUNT = 3
const DEFAULT_OUTPUT_PATH = "."

const MEM_UNIT = 1024

func MemoryFormatStrings() []string {
	return []string{"b", "kb", "mb", "gb", "tb", "pb"}
}
