package messages

type Task struct {
	IdxMin  int
	Chunksz int
	Fact    []int
	N       int
}

type Result struct {
	MaxFlips int
	CheckSum int
}
