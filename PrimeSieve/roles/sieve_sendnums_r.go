package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	sendnums2 "CodeGenTest/PrimeSieve/callbacks/result/sieve/sendnums"
	"CodeGenTest/PrimeSieve/channels/sieve/sendnums"
	"CodeGenTest/PrimeSieve/invitations"
	"sync"
)

func SieveSendNumsR(wg *sync.WaitGroup, rChan *sendnums.RChan,
	inviteChan *invitations.SieveSendNumsRInviteChan,
	env callbacks.SieveSendNumsREnv) sendnums2.RResult {
	env.Init()
SEND:
	for {
		select {
		case num := <-rChan.SNum:
			env.NumFromR(&num)
			continue SEND
		case end := <-rChan.SEnd:
			env.EndFromR(&end)
		}
		break
	}
	return env.Done()
}
