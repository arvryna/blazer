package data

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

func CalculateChunks(totalSize int, parts int) *Chunks {
	chunks := Chunks{Count: parts}
	chunks.TotalSize = totalSize
	chunks.Size = int(float64(totalSize) / float64(parts))
	pos := -1
	for i := 0; i < parts; i++ {
		r := Range{}
		r.Start = pos + 1
		pos += chunks.Size

		// Case 1
		if pos > totalSize {
			// we have already divided enough segments, so can exit early
			r.End = totalSize
			chunks.Count = i + 1
			chunks.Segments = append(chunks.Segments, r)
			break
		}

		// Case 2
		if (i == parts-1) && pos < totalSize {
			r.End = totalSize
			chunks.Segments = append(chunks.Segments, r)
			break
		}
		r.End = pos
		chunks.Segments = append(chunks.Segments, r)
	}
	return &chunks
}
