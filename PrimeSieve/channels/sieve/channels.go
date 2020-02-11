package sieve

import "CodeGenTest/PrimeSieve/messages/sieve"

type W1Chan struct {
	W2FilterPrime chan sieve.FilterPrime
}

type W2Chan struct {
	W1FilterPrime chan sieve.FilterPrime
	MPrime        chan sieve.Prime
	MFinish       chan sieve.Finish
}

type MChan struct {
	W2Prime  chan sieve.Prime
	W2Finish chan sieve.Finish
}
