package invitations

import "CodeGenTest/fannkuch-redux/channels/fannkuchrecursive"

type FannkuchRoleInviteChan struct {
	WorkerInviteMainToFannkuchRecursiveSource         chan fannkuchrecursive.SourceChan
	WorkerInviteWorkerToFannkuchRecursiveWorker       chan fannkuchrecursive.WorkerChan
	WorkerNestedInviteMainToFannkuchRecursiveSource   chan FannkuchRecursiveRoleInviteChan
	WorkerNestedInviteWorkerToFannkuchRecursiveWorker chan FannkuchRecursiveRoleInviteChan
}
