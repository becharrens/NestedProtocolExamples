package invitations

import "CodeGenTest/PrimeSieve/channels/sieve"

type PrimeSieveMasterInviteChan struct {
	WorkerInviteToSieveM           chan sieve.MChan
	WorkerInviteToSieveMInviteChan chan SieveMInviteChan
}

type PrimeSieveWorkerInviteChan struct {
	InviteWorkerToSieveW1           chan sieve.W1Chan
	InviteWorkerToSieveW1InviteChan chan SieveW1InviteChan
	InviteMasterToSieveM            chan sieve.MChan
	InviteMasterToSieveMInviteChan  chan SieveMInviteChan
}
