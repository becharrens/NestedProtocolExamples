package roles

import (
	"CodeGenTest/PrimeSieve/channels/sieve/sendnums"
	"CodeGenTest/PrimeSieve/invitations"
	sendnums2 "CodeGenTest/PrimeSieve/messages/sieve/sendnums"
	"sync"
)

func SieveSendNumsSendCommChannels(wg *sync.WaitGroup,
	inviteChan *invitations.SieveSendNumsInviteChan,
	nestedInviteChan *invitations.SieveSendNumsNestedInviteChan) {
	srNum := make(chan sendnums2.Num)
	srEnd := make(chan sendnums2.End)

	sChan := sendnums.SChan{
		RNum: srNum,
		REnd: srEnd,
	}

	rChan := sendnums.RChan{
		SNum: srNum,
		SEnd: srEnd,
	}

	rInviteChan := invitations.SieveSendNumsRInviteChan{}
	sInviteChan := invitations.SieveSendNumsSInviteChan{}

	go func() {
		inviteChan.SInviteChan <- sChan
		nestedInviteChan.SNestedInviteChan <- sInviteChan
	}()

	go func() {
		inviteChan.RInviteChan <- rChan
		nestedInviteChan.RNestedInviteChan <- rInviteChan
	}()
}
