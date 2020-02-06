package callbacks

import (
	"CodeGenTest/fannkuch-redux/callbacks/result/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/messages"
	"fmt"
)

type FannkuchRecursiveSourceEnv interface {
	Init()
	ReceiveResultFromSource(result *fannkuchrecursive.SourceResult)
	ResultFromNewWorker(result *messages.Result)
	ResultFromNewWorker2(result *messages.Result)
	Done() fannkuchrecursive.SourceResult

	ToFannkuchRecursiveSourceEnv() FannkuchRecursiveSourceEnv
}

type fannkuchRecursiveSourceState struct {
	results []*messages.Result
}

func (f *fannkuchRecursiveSourceState) Init() {
	fmt.Println("fr source started")
}

func (f *fannkuchRecursiveSourceState) ReceiveResultFromSource(result *fannkuchrecursive.SourceResult) {
	f.results = result.PartialRes
}

func (f *fannkuchRecursiveSourceState) ResultFromNewWorker(result *messages.Result) {
	f.results = append(f.results, result)
}

func (f *fannkuchRecursiveSourceState) ResultFromNewWorker2(result *messages.Result) {
	f.results = append(f.results, result)
}

func (f *fannkuchRecursiveSourceState) Done() fannkuchrecursive.SourceResult {
	fmt.Println("fr source done")
	return fannkuchrecursive.SourceResult{PartialRes: f.results}
}

func (f *fannkuchRecursiveSourceState) ToFannkuchRecursiveSourceEnv() FannkuchRecursiveSourceEnv {
	return &fannkuchRecursiveSourceState{results: f.results}
}
