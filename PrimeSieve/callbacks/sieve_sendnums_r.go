package callbacks

import (
	sendnums2 "CodeGenTest/PrimeSieve/callbacks/result/sieve/sendnums"
	"CodeGenTest/PrimeSieve/messages/sieve/sendnums"
	"fmt"
)

type SieveSendNumsREnv interface {
	Init()
	NumFromR(num *sendnums.Num)
	EndFromR(end *sendnums.End)
	Done() sendnums2.RResult
}

type sieveSendNumsRState struct {
	numsReceived []int
}

func (s *sieveSendNumsRState) Init() {
	fmt.Println("Start R")
}

func (s *sieveSendNumsRState) NumFromR(num *sendnums.Num) {
	s.numsReceived = append(s.numsReceived, num.Num)
}

func (s *sieveSendNumsRState) EndFromR(end *sendnums.End) {
}

func (s *sieveSendNumsRState) Done() sendnums2.RResult {
	return sendnums2.RResult{
		ReceivedNums: s.numsReceived,
	}
}
