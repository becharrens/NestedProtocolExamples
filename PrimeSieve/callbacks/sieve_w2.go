package callbacks

import (
	sieve2 "CodeGenTest/PrimeSieve/callbacks/result/sieve"
	"CodeGenTest/PrimeSieve/callbacks/result/sieve/sendnums"
	"CodeGenTest/PrimeSieve/messages/sieve"
	"fmt"
)

type SieveW2Choice int

const (
	PrimeToM SieveW2Choice = iota
	FinishToM
)

type SieveW2Env interface {
	Init()
	FilterPrimeFromW1(filterPrime *sieve.FilterPrime)
	AcceptSieveSendNumsRFromW1()
	ReceiveResultFromSieveSendNumsR(result *sendnums.RResult)
	W2Choice() SieveW2Choice
	PrimeToM() sieve.Prime
	PrimeSentToM()
	AcceptSieveW1FromW2()
	ReceiveResultFromSieveW1(result *sieve2.W1Result)
	FinishToM() sieve.Finish
	FinishSentToM()
	Done()

	ToSieveW1Env() SieveW1Env
	ToSieveSendNumsREnv() SieveSendNumsREnv
}

type sieveW2State struct {
	filterPrime    int
	possiblePrimes []int
}

func (s *sieveW2State) Init() {
	fmt.Println("Start W2")

}

func (s *sieveW2State) FilterPrimeFromW1(filterPrime *sieve.FilterPrime) {
	s.filterPrime = filterPrime.Prime
}

func (s *sieveW2State) AcceptSieveSendNumsRFromW1() {
}

func (s *sieveW2State) ReceiveResultFromSieveSendNumsR(result *sendnums.RResult) {
	s.possiblePrimes = result.ReceivedNums
	s.possiblePrimes = filterPrimes(s.possiblePrimes, s.filterPrime)
}

func (s *sieveW2State) W2Choice() SieveW2Choice {
	if len(s.possiblePrimes) == 0 {
		return FinishToM
	}
	return PrimeToM
}

func (s *sieveW2State) PrimeToM() sieve.Prime {
	return sieve.Prime{s.possiblePrimes[0]}
}

func (s *sieveW2State) PrimeSentToM() {
}

func (s *sieveW2State) AcceptSieveW1FromW2() {
}

func (s *sieveW2State) ReceiveResultFromSieveW1(result *sieve2.W1Result) {
}

func (s *sieveW2State) FinishToM() sieve.Finish {
	return sieve.Finish{}
}

func (s *sieveW2State) FinishSentToM() {
}

func (s *sieveW2State) Done() {
}

func (s *sieveW2State) ToSieveW1Env() SieveW1Env {
	return &sieveW1State{
		filterPrime:    s.possiblePrimes[0],
		possiblePrimes: s.possiblePrimes[1:],
	}
}

func (s *sieveW2State) ToSieveSendNumsREnv() SieveSendNumsREnv {
	return &sieveSendNumsRState{}
}

func NewSieveW2Env() SieveW2Env {
	return &sieveW2State{}
}

func filterPrimes(nums []int, prime int) []int {
	var result []int
	for _, num := range nums {
		if num%prime > 0 {
			result = append(result, num)
		}
	}
	return result
}
