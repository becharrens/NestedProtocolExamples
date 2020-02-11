package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	sieve2 "CodeGenTest/PrimeSieve/callbacks/result/sieve"
	"CodeGenTest/PrimeSieve/channels/sieve"
	"CodeGenTest/PrimeSieve/invitations"
	"sync"
)

func SieveW1(wg *sync.WaitGroup, w1Chan *sieve.W1Chan,
	inviteChan *invitations.SieveW1InviteChan, env callbacks.SieveW1Env) sieve2.W1Result {
	env.Init()

	filterPrime := env.FilterPrimeToW2()
	w1Chan.W2FilterPrime <- filterPrime
	env.FilterPrimeSentToW2()

	sieveSendNumsInviteChan := &invitations.SieveSendNumsInviteChan{
		SInviteChan: inviteChan.InviteW1ToSieveSendNumsS,
		RInviteChan: inviteChan.InviteW2ToSieveSendNumsR,
	}

	sieveSendNumsNestedInviteChan := &invitations.SieveSendNumsNestedInviteChan{
		SNestedInviteChan: inviteChan.InviteW1ToSieveSendNumsSInviteChan,
		RNestedInviteChan: inviteChan.InviteW2ToSieveSendNumsRInviteChan,
	}

	SieveSendNumsSendCommChannels(wg, sieveSendNumsInviteChan, sieveSendNumsNestedInviteChan)

	sieveSendNumsSChan := <-inviteChan.InviteW1ToSieveSendNumsS
	sieveSendNumsSInviteChan := <-inviteChan.InviteW1ToSieveSendNumsSInviteChan
	env.AcceptSieveSendNumsS()

	sieveSendNumsSEnv := env.ToSieveSendNumsSEnv()
	sieveSendNumsSResult := SieveSendNumsS(wg, &sieveSendNumsSChan,
		&sieveSendNumsSInviteChan, sieveSendNumsSEnv)
	env.ReceiveResultFromSieveSendNumsS(&sieveSendNumsSResult)

	return env.Done()
}
