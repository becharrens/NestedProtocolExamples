package invitations

import (
	"CodeGenTest/PrimeSieve/channels/sieve"
	"CodeGenTest/PrimeSieve/channels/sieve/sendnums"
)

type SieveMInviteChan struct {
	W2InviteToSieveM           chan sieve.MChan
	W2InviteToSieveMInviteChan chan SieveMInviteChan
}

type SieveW1InviteChan struct {
	InviteW1ToSieveSendNumsS           chan sendnums.SChan
	InviteW1ToSieveSendNumsSInviteChan chan SieveSendNumsSInviteChan
	InviteW2ToSieveSendNumsR           chan sendnums.RChan
	InviteW2ToSieveSendNumsRInviteChan chan SieveSendNumsRInviteChan
}

type SieveW2InviteChan struct {
	W1InviteToSieveSendNumsR           chan sendnums.RChan
	W1InviteToSieveSendNumsRInviteChan chan SieveSendNumsRInviteChan
	InviteW2ToSieveW1                  chan sieve.W1Chan
	InviteW2ToSieveW1InviteChan        chan SieveW1InviteChan
	InviteMToSieveM                    chan sieve.MChan
	InviteMToSieveMInviteChan          chan SieveMInviteChan
}

type SieveInviteChan struct {
	W1InviteChan chan sieve.W1Chan
	MInviteChan  chan sieve.MChan
}

type SieveNestedInviteChan struct {
	W1NestedInviteChan chan SieveW1InviteChan
	MNestedInviteChan  chan SieveMInviteChan
}
