package callbacks

import (
	primesieve2 "CodeGenTest/PrimeSieve/callbacks/result/primesieve"
	"CodeGenTest/PrimeSieve/callbacks/result/sieve"
	"CodeGenTest/PrimeSieve/messages/primesieve"
	"fmt"
	"strconv"
	"strings"
)

type PrimeSieveMasterEnv interface {
	Init()
	FirstPrimeToWorker() primesieve.FirstPrime
	FirstPrimeSentToWorker()
	UboundToWorker() primesieve.UBound
	UboundSentToWorker()
	AcceptSieveMFromWorker()
	ReceiveResultFromSieveM(result *sieve.MResult)
	PrimeFromWorker(prime *primesieve.Prime)
	FinishFromWorker(finish *primesieve.Finish)
	Done() primesieve2.MasterResult

	ToSieveMEnv() SieveMEnv
}

type primeSieveMasterState struct {
	n      int
	primes []int
}

func (p *primeSieveMasterState) Init() {
	fmt.Println("Start Master")
}

func (p *primeSieveMasterState) FirstPrimeToWorker() primesieve.FirstPrime {
	return primesieve.FirstPrime{
		Prime: 2,
	}
}

func (p *primeSieveMasterState) FirstPrimeSentToWorker() {
}

func (p *primeSieveMasterState) UboundToWorker() primesieve.UBound {
	return primesieve.UBound{UBound: p.n}
}

func (p *primeSieveMasterState) UboundSentToWorker() {
}

func (p *primeSieveMasterState) PrimeFromWorker(prime *primesieve.Prime) {
	p.primes = append(p.primes, prime.Prime)
}

func (p *primeSieveMasterState) AcceptSieveMFromWorker() {
}

func (p *primeSieveMasterState) ReceiveResultFromSieveM(result *sieve.MResult) {
	p.primes = result.Primes
}

func (p *primeSieveMasterState) FinishFromWorker(finish *primesieve.Finish) {
}

func (p *primeSieveMasterState) Done() primesieve2.MasterResult {
	primes := make([]string, len(p.primes))
	for i, prime := range p.primes {
		primes[i] = strconv.Itoa(prime)
	}
	fmt.Printf("Primes up to %d: %s\n", p.n, strings.Join(primes, ", "))
	return primesieve2.MasterResult{}
}

func (p *primeSieveMasterState) ToSieveMEnv() SieveMEnv {
	return &sieveMState{primes: p.primes}
}

func NewPrimeSieveMasterEnv(n int) PrimeSieveMasterEnv {
	return &primeSieveMasterState{n: n}
}
