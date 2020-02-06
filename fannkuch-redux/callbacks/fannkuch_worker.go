package callbacks

import (
	"CodeGenTest/fannkuch-redux/callbacks/result/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/messages"
	"fmt"
)

type FannkuchWorkerChoice int

const (
	CallFannkuchRecursive FannkuchWorkerChoice = iota
	ResultToMain
)

type FannkuchWorkerEnv interface {
	Init()
	TaskFromMain(task *messages.Task)
	WorkerChoice() FannkuchWorkerChoice
	ReceiveResultFromWorker(result *fannkuchrecursive.WorkerResult)
	ResultToMain() messages.Result
	SentResultToMain()
	ResultToMain2() messages.Result
	SentResultSentToMain2()
	Done()

	ToFannkuchRecursiveWorkerState() FannkuchRecursiveWorkerEnv
}

type fannkuchWorkerState struct {
	Fact    []int
	idxMin  int
	idxMax  int
	n       int
	Chunksz int
	resChan chan messages.Result
}

func (f *fannkuchWorkerState) Init() {
	f.resChan = make(chan messages.Result)
	fmt.Println("f worker started")
}

func (f *fannkuchWorkerState) TaskFromMain(task *messages.Task) {
	f.Fact = task.Fact
	f.idxMin = task.IdxMin
	f.idxMax = task.IdxMin + task.Chunksz
	f.n = task.N
	f.Chunksz = task.Chunksz
}

func (f *fannkuchWorkerState) WorkerChoice() FannkuchWorkerChoice {
	var choice FannkuchWorkerChoice
	if f.idxMax < f.Fact[f.n] {
		choice = CallFannkuchRecursive
	} else {
		f.idxMax = f.Fact[f.n]
		choice = ResultToMain
	}
	go fannkuch(f.idxMin, f.idxMax, f.n, f.Fact, f.resChan)
	fmt.Println("recurse = ", CallFannkuchRecursive == choice)
	return choice
}

func (f *fannkuchWorkerState) ReceiveResultFromWorker(result *fannkuchrecursive.WorkerResult) {
	// Do nothing. This result is just a dummy
}

func (f *fannkuchWorkerState) ResultToMain() messages.Result {
	return <-f.resChan
}

func (f *fannkuchWorkerState) SentResultToMain() {
	// Do nothing, everything is done
}

func (f *fannkuchWorkerState) ResultToMain2() messages.Result {
	return <-f.resChan
}

func (f *fannkuchWorkerState) SentResultSentToMain2() {
	// Do nothing, everything is done
}

func (f *fannkuchWorkerState) Done() {
	// Do nothing, everything is done
	fmt.Println("f worker done")
}

func (f *fannkuchWorkerState) ToFannkuchRecursiveWorkerState() FannkuchRecursiveWorkerEnv {
	return &fannkuchRecursiveWorkerState{
		Fact:    f.Fact,
		idxMin:  f.idxMax,
		n:       f.n,
		Chunksz: f.Chunksz,
	}
}

func fannkuch(idxMin int, idxMax int, n int, fact []int, resChan chan messages.Result) {
	fmt.Printf("fannkuch called with %d, %d\n", idxMin, idxMax)

	p := make([]int, n)
	pp := make([]int, n)
	count := make([]int, n)

	// first permutation
	for i := 0; i < n; i++ {
		p[i] = i
	}
	for i, idx := n-1, idxMin; i > 0; i-- {
		d := idx / fact[i]
		count[i] = d
		idx = idx % fact[i]

		copy(pp, p)
		for j := 0; j <= i; j++ {
			if j+d <= i {
				p[j] = pp[j+d]
			} else {
				p[j] = pp[j+d-i-1]
			}
		}
	}

	maxFlips := 1
	checkSum := 0

	for idx, sign := idxMin, true; ; sign = !sign {

		// count flips
		first := p[0]
		if first != 0 {
			flips := 1
			if p[first] != 0 {
				copy(pp, p)
				p0 := first
				for {
					flips++
					for i, j := 1, p0-1; i < j; i, j = i+1, j-1 {
						pp[i], pp[j] = pp[j], pp[i]
					}
					t := pp[p0]
					pp[p0] = p0
					p0 = t
					if pp[p0] == 0 {
						break
					}
				}
			}
			if maxFlips < flips {
				maxFlips = flips
			}
			if sign {
				checkSum += flips
			} else {
				checkSum -= flips
			}
		}

		if idx++; idx == idxMax {
			break
		}

		// next permutation
		if sign {
			p[0], p[1] = p[1], first
		} else {
			p[1], p[2] = p[2], p[1]
			for k := 2; ; k++ {
				if count[k]++; count[k] <= k {
					break
				}
				count[k] = 0
				for j := 0; j <= k; j++ {
					p[j] = p[j+1]
				}
				p[k+1] = first
				first = p[0]
			}
		}
	}

	resChan <- messages.Result{MaxFlips: maxFlips, CheckSum: checkSum}
}

func NewFannkuchWorkerEnv() FannkuchWorkerEnv {
	return &fannkuchWorkerState{}
}
