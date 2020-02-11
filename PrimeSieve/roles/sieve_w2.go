package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	"CodeGenTest/PrimeSieve/channels/sieve"
	"CodeGenTest/PrimeSieve/invitations"
	"sync"
)

func SieveW2(wg *sync.WaitGroup, w2Chan *sieve.W2Chan,
	inviteChan *invitations.SieveW2InviteChan, env callbacks.SieveW2Env) {
	defer wg.Done()
	env.Init()

	filterPrime := <-w2Chan.W1FilterPrime
	env.FilterPrimeFromW1(&filterPrime)

	sieveSendNumsRChan := <-inviteChan.W1InviteToSieveSendNumsR
	sieveSendNumsRInviteChan := <-inviteChan.W1InviteToSieveSendNumsRInviteChan
	env.AcceptSieveSendNumsRFromW1()

	sieveSendNumsREnv := env.ToSieveSendNumsREnv()
	sieveSendNumsRResult := SieveSendNumsR(wg, &sieveSendNumsRChan,
		&sieveSendNumsRInviteChan, sieveSendNumsREnv)
	env.ReceiveResultFromSieveSendNumsR(&sieveSendNumsRResult)

	w2Choice := env.W2Choice()
	switch w2Choice {
	case callbacks.PrimeToM:
		prime := env.PrimeToM()
		w2Chan.MPrime <- prime
		env.PrimeSentToM()

		sieveInviteChan := &invitations.SieveInviteChan{
			W1InviteChan: inviteChan.InviteW2ToSieveW1,
			MInviteChan:  inviteChan.InviteMToSieveM,
		}
		sieveNestedInviteChan := &invitations.SieveNestedInviteChan{
			W1NestedInviteChan: inviteChan.InviteW2ToSieveW1InviteChan,
			MNestedInviteChan:  inviteChan.InviteMToSieveMInviteChan,
		}
		SieveSendCommChannels(wg, sieveInviteChan, sieveNestedInviteChan)

		sieveW1Chan := <-inviteChan.InviteW2ToSieveW1
		sieveW1InviteChan := <-inviteChan.InviteW2ToSieveW1InviteChan
		env.AcceptSieveW1FromW2()

		sieveW1Env := env.ToSieveW1Env()
		sieveW1Result := SieveW1(wg, &sieveW1Chan, &sieveW1InviteChan, sieveW1Env)
		env.ReceiveResultFromSieveW1(&sieveW1Result)

	case callbacks.FinishToM:
		finish := env.FinishToM()
		w2Chan.MFinish <- finish
		env.FinishSentToM()
	}

	env.Done()
}
