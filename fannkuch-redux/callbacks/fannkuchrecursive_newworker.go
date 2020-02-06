package callbacks

import (
	"CodeGenTest/fannkuch-redux/callbacks/result/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/messages"
	"fmt"
)

type FannkuchRecursiveNewWorkerChoice int

const (
	CallFannkuchRecursive2 FannkuchRecursiveNewWorkerChoice = iota
	ResultToSource
)

type FannkuchRecursiveNewWorkerEnv interface {
	Init()
	TaskFromSource(task *messages.Task)
	NewWorkerChoice() FannkuchRecursiveNewWorkerChoice
	ReceiveResultFromWorker(result *fannkuchrecursive.WorkerResult)
	ResultToSource() messages.Result
	ResultSentToSource()
	ResultToSource2() messages.Result
	ResultSentToSource2()
	Done()

	ToFannkuchRecursiveWorkerEnv() FannkuchRecursiveWorkerEnv
}

type fannkuchRecursiveNewWorkerState struct {
	Fact    []int
	idxMin  int
	idxMax  int
	n       int
	resChan chan messages.Result
	Chunksz int
}

func (f *fannkuchRecursiveNewWorkerState) Init() {
	f.resChan = make(chan messages.Result)
	fmt.Println("fr new worker started")
}

func (f *fannkuchRecursiveNewWorkerState) TaskFromSource(task *messages.Task) {
	f.Fact = task.Fact
	f.idxMin = task.IdxMin
	f.idxMax = task.IdxMin + task.Chunksz
	f.n = task.N
	f.Chunksz = task.Chunksz
}

func (f *fannkuchRecursiveNewWorkerState) NewWorkerChoice() FannkuchRecursiveNewWorkerChoice {
	var choice FannkuchRecursiveNewWorkerChoice
	if f.idxMax < f.Fact[f.n] {
		choice = CallFannkuchRecursive2
	} else {
		f.idxMax = f.Fact[f.n]
		choice = ResultToSource
	}
	go fannkuch(f.idxMin, f.idxMax, f.n, f.Fact, f.resChan)
	fmt.Println("recurse = ", CallFannkuchRecursive2 == choice)
	return choice
}

func (f *fannkuchRecursiveNewWorkerState) ReceiveResultFromWorker(result *fannkuchrecursive.WorkerResult) {
	// Do nothing
}

func (f *fannkuchRecursiveNewWorkerState) ResultToSource() messages.Result {
	return <-f.resChan
}

func (f *fannkuchRecursiveNewWorkerState) ResultSentToSource() {
	// Do nothing
}

func (f *fannkuchRecursiveNewWorkerState) ResultToSource2() messages.Result {
	return <-f.resChan
}

func (f *fannkuchRecursiveNewWorkerState) ResultSentToSource2() {
	// Do nothing
}

func (f *fannkuchRecursiveNewWorkerState) Done() {
	fmt.Println("fr new worker done")
}

func (f *fannkuchRecursiveNewWorkerState) ToFannkuchRecursiveWorkerEnv() FannkuchRecursiveWorkerEnv {
	return &fannkuchRecursiveWorkerState{
		Fact:    f.Fact,
		idxMin:  f.idxMax,
		n:       f.n,
		Chunksz: f.Chunksz,
	}
}

func NewFannkuchRecursiveNewWorkerEnv() FannkuchRecursiveNewWorkerEnv {
	return &fannkuchRecursiveNewWorkerState{}
}
