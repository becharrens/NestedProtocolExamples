package roles

import (
	"CodeGenTest/PrimeSieve/callbacks"
	"CodeGenTest/PrimeSieve/channels/sieve"
	"CodeGenTest/PrimeSieve/channels/sieve/sendnums"
	"CodeGenTest/PrimeSieve/invitations"
	sieve2 "CodeGenTest/PrimeSieve/messages/sieve"
	"sync"
)

func SieveSendCommChannels(wg *sync.WaitGroup,
	inviteChan *invitations.SieveInviteChan,
	nestedInviteChan *invitations.SieveNestedInviteChan) {

	mW2Prime := make(chan sieve2.Prime)
	mW2Finish := make(chan sieve2.Finish)
	w1W2FilterPrime := make(chan sieve2.FilterPrime)

	mChan := sieve.MChan{
		W2Prime:  mW2Prime,
		W2Finish: mW2Finish,
	}

	w1Chan := sieve.W1Chan{
		W2FilterPrime: w1W2FilterPrime,
	}

	w2Chan := sieve.W2Chan{
		W1FilterPrime: w1W2FilterPrime,
		MPrime:        mW2Prime,
		MFinish:       mW2Finish,
	}

	w2InviteMToSieveM := make(chan sieve.MChan)
	w2InviteMToSieveMInviteChan := make(chan invitations.SieveMInviteChan)
	w2InviteW2ToSieveW1 := make(chan sieve.W1Chan)
	w2InviteW2ToSieveW1InviteChan := make(chan invitations.SieveW1InviteChan)

	w1InviteW1ToSieveSendNumsS := make(chan sendnums.SChan)
	w1InviteW1ToSieveSendNumsSInviteChan := make(chan invitations.SieveSendNumsSInviteChan)
	w1InviteW2ToSieveSendNumsR := make(chan sendnums.RChan)
	w1InviteW2ToSieveSendNumsRInviteChan := make(chan invitations.SieveSendNumsRInviteChan)

	mInviteChan := invitations.SieveMInviteChan{
		W2InviteToSieveM:           w2InviteMToSieveM,
		W2InviteToSieveMInviteChan: w2InviteMToSieveMInviteChan,
	}

	w1InviteChan := invitations.SieveW1InviteChan{
		InviteW1ToSieveSendNumsS:           w1InviteW1ToSieveSendNumsS,
		InviteW1ToSieveSendNumsSInviteChan: w1InviteW1ToSieveSendNumsSInviteChan,
		InviteW2ToSieveSendNumsR:           w1InviteW2ToSieveSendNumsR,
		InviteW2ToSieveSendNumsRInviteChan: w1InviteW2ToSieveSendNumsRInviteChan,
	}

	w2InviteChan := invitations.SieveW2InviteChan{
		W1InviteToSieveSendNumsR:           w1InviteW2ToSieveSendNumsR,
		W1InviteToSieveSendNumsRInviteChan: w1InviteW2ToSieveSendNumsRInviteChan,
		InviteW2ToSieveW1:                  w2InviteW2ToSieveW1,
		InviteW2ToSieveW1InviteChan:        w2InviteW2ToSieveW1InviteChan,
		InviteMToSieveM:                    w2InviteMToSieveM,
		InviteMToSieveMInviteChan:          w2InviteMToSieveMInviteChan,
	}

	go func() {
		inviteChan.MInviteChan <- mChan
		nestedInviteChan.MNestedInviteChan <- mInviteChan
	}()

	go func() {
		inviteChan.W1InviteChan <- w1Chan
		nestedInviteChan.W1NestedInviteChan <- w1InviteChan
	}()

	wg.Add(1)
	w2Env := callbacks.NewSieveW2Env()
	go SieveW2(wg, &w2Chan, &w2InviteChan, w2Env)
}
