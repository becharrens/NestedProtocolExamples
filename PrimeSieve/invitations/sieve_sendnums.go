package invitations

import "CodeGenTest/PrimeSieve/channels/sieve/sendnums"

type SieveSendNumsSInviteChan struct {
}

type SieveSendNumsRInviteChan struct {
}

type SieveSendNumsInviteChan struct {
	SInviteChan chan sendnums.SChan
	RInviteChan chan sendnums.RChan
}

type SieveSendNumsNestedInviteChan struct {
	SNestedInviteChan chan SieveSendNumsSInviteChan
	RNestedInviteChan chan SieveSendNumsRInviteChan
}
