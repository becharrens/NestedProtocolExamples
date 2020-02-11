package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	primesieve2 "CodeGenTest/PrimeSieve/callbacks/result/primesieve"
	"CodeGenTest/PrimeSieve/channels/primesieve"
	"CodeGenTest/PrimeSieve/invitations"
	"sync"
)

func StartPrimeSieveWorker(wg *sync.WaitGroup, workerChan *primesieve.WorkerChan,
	inviteChan *invitations.PrimeSieveWorkerInviteChan,
	env callbacks.PrimeSieveWorkerEnv) {
	defer wg.Done()
	PrimeSieveWorker(wg, workerChan, inviteChan, env)
}

func PrimeSieveWorker(wg *sync.WaitGroup, workerChan *primesieve.WorkerChan,
	inviteChan *invitations.PrimeSieveWorkerInviteChan,
	env callbacks.PrimeSieveWorkerEnv) primesieve2.WorkerResult {
	env.Init()

	firstPrime := <-workerChan.MasterFirstPrime
	env.FirstPrimeFromMaster(&firstPrime)

	ubound := <-workerChan.MasterUBound
	env.UboundFromMaster(&ubound)

	workerChoice := env.WorkerChoice()
	switch workerChoice {
	case callbacks.PrimeToMster:
		prime := env.PrimeToMaster()
		workerChan.MasterPrime <- prime
		env.PrimeSentToMaster()

		sieveInviteChan := &invitations.SieveInviteChan{
			W1InviteChan: inviteChan.InviteWorkerToSieveW1,
			MInviteChan:  inviteChan.InviteMasterToSieveM,
		}
		sieveNestedInviteChan := &invitations.SieveNestedInviteChan{
			W1NestedInviteChan: inviteChan.InviteWorkerToSieveW1InviteChan,
			MNestedInviteChan:  inviteChan.InviteMasterToSieveMInviteChan,
		}
		SieveSendCommChannels(wg, sieveInviteChan, sieveNestedInviteChan)

		sieveW1Chan := <-inviteChan.InviteWorkerToSieveW1
		sieveW1InviteChan := <-inviteChan.InviteWorkerToSieveW1InviteChan
		env.AcceptSieveW1FromWorker()

		sieveW1Env := env.ToSieveW1Env()
		sieveW1result := SieveW1(wg, &sieveW1Chan, &sieveW1InviteChan, sieveW1Env)
		env.ReceiveResultFromSieveW1(&sieveW1result)

	case callbacks.FinishToMaster:
		finish := env.FinishToMaster()
		workerChan.MasterFinish <- finish
		env.FinishSentToMaster()
	}

	return env.Done()
}
