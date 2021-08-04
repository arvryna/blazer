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

func (c *Chunks) ComputeChunks() {
	c.Size = int(float64(c.TotalSize) / float64(c.Count))
	pos := -1
	for i := 0; i < c.Count; i++ {
		r := Range{}
		r.Start = pos + 1
		pos += c.Size

		// Case 1
		if pos > c.TotalSize {
			// we have already divided enough segments, so can exit early
			r.End = c.TotalSize
			c.Count = i + 1
			c.Segments = append(c.Segments, r)
			break
		}

		// Case 2
		if (i == c.Count-1) && pos < c.TotalSize {
			r.End = c.TotalSize
			c.Segments = append(c.Segments, r)
			break
		}
		r.End = pos
		c.Segments = append(c.Segments, r)
	}
}
