package data

type Range struct {
	Start int
	End   int
}

type Chunks struct {
	ChunkSize int
	TotalSize int
	Segments  []Range
	Parts     int
}

func CalculateChunks(totalSize int, parts int) *Chunks {
	chunks := Chunks{}
	val := float64(totalSize) / float64(parts)
	// There is a bug here that last segment will get the largest chunk , that needs to be fixed

	// Trade of, to force flooring or force math.Round, less thread or more thread ????
	// if val > math.Round(val) {
	// 	val = math.Round(val) + 1
	// }
	// update the return type values properly other wise you will fill wrong values
	// refractor this code
	// Handle if the requested thread size is just 1
	// later compare this code with other sources, also learn more about it

	chunkSize := int(val)
	chunks.TotalSize = totalSize
	chunks.Parts = parts
	chunks.ChunkSize = chunkSize
	pos := -1
	for i := 0; i < parts; i++ {
		r := Range{}
		r.Start = pos + 1
		pos += chunkSize

		// Case 1
		if pos > totalSize {
			// we have already divided enough segments, so can exit early
			r.End = totalSize
			chunks.Parts = i + 1
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
