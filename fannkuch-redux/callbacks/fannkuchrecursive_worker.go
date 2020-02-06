package callbacks

import (
	"CodeGenTest/fannkuch-redux/callbacks/result/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/messages"
	"fmt"
)

type FannkuchRecursiveWorkerEnv interface {
	Init()
	TaskToNewWorker() messages.Task
	SentTaskToNewWorker()
	Done() fannkuchrecursive.WorkerResult
}

type fannkuchRecursiveWorkerState struct {
	Fact    []int
	idxMin  int
	n       int
	Chunksz int
}

func (f *fannkuchRecursiveWorkerState) Init() {
	fmt.Println("fr worker started")
}

func (f *fannkuchRecursiveWorkerState) TaskToNewWorker() messages.Task {
	return messages.Task{
		IdxMin:  f.idxMin,
		Chunksz: f.Chunksz,
		Fact:    f.Fact,
		N:       f.n,
	}
}

func (f *fannkuchRecursiveWorkerState) SentTaskToNewWorker() {
	// Do nothing
}

func (f *fannkuchRecursiveWorkerState) Done() fannkuchrecursive.WorkerResult {
	fmt.Println("fr worker done")
	return fannkuchrecursive.WorkerResult{}
}
