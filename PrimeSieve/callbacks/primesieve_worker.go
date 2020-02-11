package callbacks

import (
	primesieve2 "CodeGenTest/PrimeSieve/callbacks/result/primesieve"
	"CodeGenTest/PrimeSieve/callbacks/result/sieve"
	"CodeGenTest/PrimeSieve/messages/primesieve"
	"fmt"
)

type PrimeSieveWorkerChoice int

const (
	PrimeToMster PrimeSieveWorkerChoice = iota
	FinishToMaster
)

type PrimeSieveWorkerEnv interface {
	Init()
	FirstPrimeFromMaster(prime *primesieve.FirstPrime)
	UboundFromMaster(prime *primesieve.UBound)
	WorkerChoice() PrimeSieveWorkerChoice
	AcceptSieveW1FromWorker()
	ReceiveResultFromSieveW1(result *sieve.W1Result)
	PrimeToMaster() primesieve.Prime
	PrimeSentToMaster()
	FinishToMaster() primesieve.Finish
	FinishSentToMaster()
	Done() primesieve2.WorkerResult

	ToSieveW1Env() SieveW1Env
}

type primeSieveWorkerState struct {
	firstPrime     int
	uBound         int
	possiblePrimes []int
}

func (p *primeSieveWorkerState) Init() {
	fmt.Println("Start Worker")
}

func (p *primeSieveWorkerState) FirstPrimeFromMaster(prime *primesieve.FirstPrime) {
	p.firstPrime = prime.Prime
}

func (p *primeSieveWorkerState) UboundFromMaster(uBound *primesieve.UBound) {
	p.uBound = uBound.UBound
	p.possiblePrimes = initPossiblePrimes(p.firstPrime, p.uBound)
}

func (p *primeSieveWorkerState) WorkerChoice() PrimeSieveWorkerChoice {
	if len(p.possiblePrimes) == 0 {
		return FinishToMaster
	}
	return PrimeToMster
}

func (p *primeSieveWorkerState) PrimeToMaster() primesieve.Prime {
	return primesieve.Prime{Prime: p.possiblePrimes[0]}
}

func (p *primeSieveWorkerState) PrimeSentToMaster() {
}

func (p *primeSieveWorkerState) AcceptSieveW1FromWorker() {
}

func (p *primeSieveWorkerState) ReceiveResultFromSieveW1(result *sieve.W1Result) {
}

func (p *primeSieveWorkerState) FinishToMaster() primesieve.Finish {
	return primesieve.Finish{}
}

func (p *primeSieveWorkerState) FinishSentToMaster() {
}

func (p *primeSieveWorkerState) Done() primesieve2.WorkerResult {
	return primesieve2.WorkerResult{}
}

func (p *primeSieveWorkerState) ToSieveW1Env() SieveW1Env {
	return &sieveW1State{
		filterPrime:    p.possiblePrimes[0],
		possiblePrimes: p.possiblePrimes[1:],
	}
}

func NewPrimeSieveWorkerEnv() PrimeSieveWorkerEnv {
	return &primeSieveWorkerState{}
}

func initPossiblePrimes(firstPrime, ubound int) []int {
	var result []int
	for i := 2; i < ubound; i++ {
		if i%firstPrime > 0 {
			result = append(result, i)
		}
	}
	return result
}
