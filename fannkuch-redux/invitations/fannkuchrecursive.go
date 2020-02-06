package invitations

import "CodeGenTest/fannkuch-redux/channels/fannkuchrecursive"

type FannkuchRecursiveInviteChan struct {
	SourceInviteChan chan fannkuchrecursive.SourceChan
	WorkerInviteChan chan fannkuchrecursive.WorkerChan
}

type FannkuchRecursiveNestedInviteChan struct {
	SourceNestedInviteChan chan FannkuchRecursiveRoleInviteChan
	WorkerNestedInviteChan chan FannkuchRecursiveRoleInviteChan
}

type FannkuchRecursiveRoleInviteChan struct {
	NewWorkerInviteSourceToFannkuchRecursiveSource          chan fannkuchrecursive.SourceChan
	NewWorkerInviteNewWorkerToFannkuchRecursiveWorker       chan fannkuchrecursive.WorkerChan
	NewWorkerNestedInviteSourceToFannkuchRecursiveSource    chan FannkuchRecursiveRoleInviteChan
	NewWorkerNestedInviteNewWorkerToFannkuchRecursiveWorker chan FannkuchRecursiveRoleInviteChan
}
