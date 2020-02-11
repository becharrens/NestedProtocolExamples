package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	sendnums2 "CodeGenTest/PrimeSieve/callbacks/result/sieve/sendnums"
	"CodeGenTest/PrimeSieve/channels/sieve/sendnums"
	"CodeGenTest/PrimeSieve/invitations"
	"sync"
)

func SieveSendNumsS(wg *sync.WaitGroup, sChan *sendnums.SChan,
	inviteChan *invitations.SieveSendNumsSInviteChan,
	env callbacks.SieveSendNumsSEnv) sendnums2.SResult {
	env.Init()
SEND:
	for {
		sChoice := env.SChoice()
		switch sChoice {
		case callbacks.NumToR:
			num := env.NumToR()
			sChan.RNum <- num
			env.NumSentToR()
			continue SEND
		case callbacks.EndToR:
			end := env.EndToR()
			sChan.REnd <- end
			env.EndSentToR()
		}
		break
	}
	return env.Done()
}
