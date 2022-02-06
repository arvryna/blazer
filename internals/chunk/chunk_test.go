package chunk

import "testing"

func TestComputeChunk(t *testing.T) {
	chunks := ChunkList{TotalSize: 100, Count: 10}
	chunks.ComputeChunks()
	if len(chunks.Segments) != 10 {
		t.Error("Segment computation failed", len(chunks.Segments))
	}
}
