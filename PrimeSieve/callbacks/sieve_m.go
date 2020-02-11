package callbacks

import (
	"CodeGenTest/PrimeSieve/callbacks/result/sieve"
	sieve2 "CodeGenTest/PrimeSieve/messages/sieve"
	"fmt"
)

type SieveMEnv interface {
	Init()
	PrimeFromW2(prime *sieve2.Prime)
	AcceptSieveMFromW2()
	ReceiveResultFromSieveM(result *sieve.MResult)
	FinishFromW2(finish *sieve2.Finish)
	Done() sieve.MResult

	ToSieveMEnv() SieveMEnv
}

type sieveMState struct {
	primes []int
}

func (s *sieveMState) Init() {
	fmt.Println("Start M")
}

func (s *sieveMState) PrimeFromW2(prime *sieve2.Prime) {
	s.primes = append(s.primes, prime.Prime)
}

func (s *sieveMState) AcceptSieveMFromW2() {
}

func (s *sieveMState) ReceiveResultFromSieveM(result *sieve.MResult) {
	s.primes = result.Primes
}

func (s *sieveMState) FinishFromW2(finish *sieve2.Finish) {
}

func (s *sieveMState) Done() sieve.MResult {
	return sieve.MResult{Primes: s.primes}
}

func (s *sieveMState) ToSieveMEnv() SieveMEnv {
	return s
}
