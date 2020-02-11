package callbacks

import (
	sendnums2 "CodeGenTest/PrimeSieve/callbacks/result/sieve/sendnums"
	"CodeGenTest/PrimeSieve/messages/sieve/sendnums"
	"fmt"
)

type SieveSendNumsSChoice int

const (
	NumToR = iota
	EndToR = iota
)

type SieveSendNumsSEnv interface {
	Init()
	SChoice() SieveSendNumsSChoice
	NumToR() sendnums.Num
	NumSentToR()
	EndToR() sendnums.End
	EndSentToR()
	Done() sendnums2.SResult
}

type sieveSendNumsSState struct {
	idx        int
	numsToSend []int
}

func (s *sieveSendNumsSState) Init() {
	fmt.Println("Start S")
	s.idx = 0
}

func (s *sieveSendNumsSState) SChoice() SieveSendNumsSChoice {
	if s.idx >= len(s.numsToSend) {
		return EndToR
	}
	return NumToR
}

func (s *sieveSendNumsSState) NumToR() sendnums.Num {
	result := s.numsToSend[s.idx]
	s.idx++
	return sendnums.Num{Num: result}
}

func (s *sieveSendNumsSState) NumSentToR() {
}

func (s *sieveSendNumsSState) EndToR() sendnums.End {
	return sendnums.End{}
}

func (s *sieveSendNumsSState) EndSentToR() {
}

func (s *sieveSendNumsSState) Done() sendnums2.SResult {
	return sendnums2.SResult{}
}
