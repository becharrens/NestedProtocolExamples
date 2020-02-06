package roles

import (
	"CodeGenTest/fannkuch-redux/callbacks"
	"CodeGenTest/fannkuch-redux/channels/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/invitations"
	"CodeGenTest/fannkuch-redux/messages"
	"sync"
)

func FannkuchRecursiveSendCommChannels(wg *sync.WaitGroup,
	inviteChan *invitations.FannkuchRecursiveInviteChan,
	nestedInviteChan *invitations.FannkuchRecursiveNestedInviteChan) {

	sourceNewWorkerResult1 := make(chan messages.Result)
	sourceNewWorkerResult2 := make(chan messages.Result)
	workerNewWorkerTask := make(chan messages.Task)

	sourceChan := fannkuchrecursive.SourceChan{
		NewWorkerResult1: sourceNewWorkerResult1,
		NewWorkerResult2: sourceNewWorkerResult2,
	}
	newWorkerChan := fannkuchrecursive.NewWorkerChan{
		SourceResult1: sourceNewWorkerResult1,
		SourceResult2: sourceNewWorkerResult2,
		WorkerTask:    workerNewWorkerTask,
	}
	workerChan := fannkuchrecursive.WorkerChan{
		NewWorkerTask: workerNewWorkerTask,
	}

	roleInviteChan := invitations.FannkuchRecursiveRoleInviteChan{
		NewWorkerInviteSourceToFannkuchRecursiveSource:          make(chan fannkuchrecursive.SourceChan),
		NewWorkerInviteNewWorkerToFannkuchRecursiveWorker:       make(chan fannkuchrecursive.WorkerChan),
		NewWorkerNestedInviteSourceToFannkuchRecursiveSource:    make(chan invitations.FannkuchRecursiveRoleInviteChan),
		NewWorkerNestedInviteNewWorkerToFannkuchRecursiveWorker: make(chan invitations.FannkuchRecursiveRoleInviteChan),
	}

	go func() {
		inviteChan.SourceInviteChan <- sourceChan
		nestedInviteChan.SourceNestedInviteChan <- roleInviteChan
	}()

	go func() {
		inviteChan.WorkerInviteChan <- workerChan
		nestedInviteChan.WorkerNestedInviteChan <- roleInviteChan
	}()

	wg.Add(1)
	newWorkerEnv := callbacks.NewFannkuchRecursiveNewWorkerEnv()
	go FannkuchRecursiveNewWorker(wg, &newWorkerChan, &roleInviteChan, newWorkerEnv)
}
