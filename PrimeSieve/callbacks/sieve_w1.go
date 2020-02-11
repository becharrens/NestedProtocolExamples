package callbacks

import (
	sieve2 "CodeGenTest/PrimeSieve/callbacks/result/sieve"
	"CodeGenTest/PrimeSieve/callbacks/result/sieve/sendnums"
	"CodeGenTest/PrimeSieve/messages/sieve"
	"fmt"
)

type SieveW1Env interface {
	Init()
	FilterPrimeToW2() sieve.FilterPrime
	FilterPrimeSentToW2()
	AcceptSieveSendNumsS()
	ReceiveResultFromSieveSendNumsS(result *sendnums.SResult)
	Done() sieve2.W1Result

	ToSieveSendNumsSEnv() SieveSendNumsSEnv
}

type sieveW1State struct {
	filterPrime    int
	possiblePrimes []int
}

func (s *sieveW1State) Init() {
	fmt.Println("Start W1")
}

func (s *sieveW1State) FilterPrimeToW2() sieve.FilterPrime {
	return sieve.FilterPrime{
		Prime: s.filterPrime,
	}
}

func (s *sieveW1State) FilterPrimeSentToW2() {
}

func (s *sieveW1State) AcceptSieveSendNumsS() {
}

func (s *sieveW1State) ReceiveResultFromSieveSendNumsS(result *sendnums.SResult) {
}

func (s *sieveW1State) Done() sieve2.W1Result {
	return sieve2.W1Result{}
}

func (s *sieveW1State) ToSieveSendNumsSEnv() SieveSendNumsSEnv {
	return &sieveSendNumsSState{
		idx:        0,
		numsToSend: s.possiblePrimes,
	}
}
