package callbacks

import (
	"CodeGenTest/fannkuch-redux/callbacks/result/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/messages"
	"fmt"
)

type FannkuchMainEnv interface {
	Init()
	TaskToWorker() messages.Task
	SentTaskToWorker()
	ReceiveResultFromSource(result *fannkuchrecursive.SourceResult)
	ResultFromWorker(result *messages.Result)
	ResultFromWorker2(result *messages.Result)
	Done()

	ToFannkuchRecursiveSourceState() FannkuchRecursiveSourceEnv
}

const NCHUNKS = 720

type fannkuchMainState struct {
	// Fill in state of Main here
	Fact       []int
	partialRes []*messages.Result
	n          int
	Chunksz    int
}

func (f *fannkuchMainState) Init() {
	fmt.Println("Start Main")
}

func (f *fannkuchMainState) TaskToWorker() messages.Task {
	return messages.Task{
		IdxMin:  0,
		Chunksz: f.Chunksz,
		Fact:    f.Fact,
		N:       f.n,
	}
}

func (f *fannkuchMainState) SentTaskToWorker() {
	// Done?
}

func (f *fannkuchMainState) ReceiveResultFromSource(result *fannkuchrecursive.SourceResult) {
	f.partialRes = result.PartialRes
}

func (f *fannkuchMainState) ResultFromWorker(result *messages.Result) {
	f.partialRes = append(f.partialRes, result)
}

func (f *fannkuchMainState) ResultFromWorker2(result *messages.Result) {
	f.partialRes = append(f.partialRes, result)
}

func (f *fannkuchMainState) ToFannkuchRecursiveSourceState() FannkuchRecursiveSourceEnv {
	return &fannkuchRecursiveSourceState{results: f.partialRes}
}

func (f *fannkuchMainState) Done() {
	res := 0
	chk := 0
	for _, r := range f.partialRes {
		if res < r.MaxFlips {
			res = r.MaxFlips
		}
		chk += r.CheckSum
	}
	printResult(f.n, res, chk)
	fmt.Println("Main Done")
}

func printResult(n int, res int, chk int) {
	fmt.Printf("%d\nPfannkuchen(%d) = %d\n", chk, n, res)
}

func NewFannkuchMainEnv(n int) FannkuchMainEnv {
	f := &fannkuchMainState{
		Fact:       make([]int, n+1),
		partialRes: nil,
		n:          n,
	}

	f.Fact[0] = 1
	for i := 1; i < len(f.Fact); i++ {
		f.Fact[i] = f.Fact[i-1] * i
	}

	chunksz := (f.Fact[n] + NCHUNKS - 1) / NCHUNKS
	chunksz += chunksz % 2
	f.Chunksz = chunksz

	return f
}
