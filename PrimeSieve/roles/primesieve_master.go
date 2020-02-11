package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	primesieve2 "CodeGenTest/PrimeSieve/callbacks/result/primesieve"
	"CodeGenTest/PrimeSieve/channels/primesieve"
	"CodeGenTest/PrimeSieve/invitations"
	"sync"
)

func StartPrimeSieveMaster(wg *sync.WaitGroup, masterChan *primesieve.MasterChan,
	inviteChan *invitations.PrimeSieveMasterInviteChan,
	env callbacks.PrimeSieveMasterEnv) {
	defer wg.Done()
	PrimeSieveMaster(wg, masterChan, inviteChan, env)
}

func PrimeSieveMaster(wg *sync.WaitGroup, masterChan *primesieve.MasterChan,
	inviteChan *invitations.PrimeSieveMasterInviteChan,
	env callbacks.PrimeSieveMasterEnv) primesieve2.MasterResult {
	env.Init()

	firstPrime := env.FirstPrimeToWorker()
	masterChan.WorkerFirstPrime <- firstPrime
	env.FirstPrimeSentToWorker()

	ubound := env.UboundToWorker()
	masterChan.WorkerUBound <- ubound
	env.UboundSentToWorker()

	select {
	case prime := <-masterChan.WorkerPrime:
		env.PrimeFromWorker(&prime)

		sieveMChan := <-inviteChan.WorkerInviteToSieveM
		sieveMInviteChan := <-inviteChan.WorkerInviteToSieveMInviteChan
		env.AcceptSieveMFromWorker()

		sieveMEnv := env.ToSieveMEnv()
		sieveMresult := SieveM(wg, &sieveMChan, &sieveMInviteChan, sieveMEnv)
		env.ReceiveResultFromSieveM(&sieveMresult)

	case finish := <-masterChan.WorkerFinish:
		env.FinishFromWorker(&finish)
	}

	return env.Done()
}
