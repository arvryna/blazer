package data

type ChunkRange struct {
	Start int
	End   int
}

type Chunks struct {
	ChunkSize     int
	TotalSize     int
	Ranges        []ChunkRange
	NunberOfParts int
}

func CalculateChunks(totalSize int, partsToDivide int) *Chunks {
	chunks := Chunks{}
	chunkSize := totalSize / partsToDivide
	chunks.TotalSize = totalSize
	chunks.NunberOfParts = partsToDivide
	chunks.ChunkSize = chunkSize
	pos := -1
	for i := 0; i < partsToDivide; i++ {
		r := ChunkRange{}
		r.Start = pos + 1
		r.End = (pos + chunkSize) - 1
		pos = pos + chunkSize
		chunks.Ranges = append(chunks.Ranges, r)
	}
	return &chunks
}
