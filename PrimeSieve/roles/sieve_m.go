package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	sieve2 "CodeGenTest/PrimeSieve/callbacks/result/sieve"
	"CodeGenTest/PrimeSieve/channels/sieve"
	"CodeGenTest/PrimeSieve/invitations"
	"sync"
)

func SieveM(wg *sync.WaitGroup, mChan *sieve.MChan,
	inviteChan *invitations.SieveMInviteChan, env callbacks.SieveMEnv) sieve2.MResult {
	env.Init()
	select {
	case prime := <-mChan.W2Prime:
		env.PrimeFromW2(&prime)

		sieveMChan := <-inviteChan.W2InviteToSieveM
		sieveMInviteChan := <-inviteChan.W2InviteToSieveMInviteChan
		env.AcceptSieveMFromW2()

		sieveMEnv := env.ToSieveMEnv()
		result := SieveM(wg, &sieveMChan, &sieveMInviteChan, sieveMEnv)
		env.ReceiveResultFromSieveM(&result)
	case finish := <-mChan.W2Finish:
		env.FinishFromW2(&finish)
	}
	return env.Done()
}
