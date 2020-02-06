package roles

import (
	"CodeGenTest/fannkuch-redux/callbacks"
	"CodeGenTest/fannkuch-redux/channels/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/invitations"
	"fmt"
	"sync"
)

func FannkuchRecursiveNewWorker(wg *sync.WaitGroup,
	newWorkerChan *fannkuchrecursive.NewWorkerChan,
	inviteChan *invitations.FannkuchRecursiveRoleInviteChan,
	env callbacks.FannkuchRecursiveNewWorkerEnv) {
	defer wg.Done()

	env.Init()

	task := <-newWorkerChan.WorkerTask
	env.TaskFromSource(&task)

	choice := env.NewWorkerChoice()
	switch choice {
	case callbacks.CallFannkuchRecursive2:
		fmt.Println("NewWorker recurse")
		inviteChannels := &invitations.FannkuchRecursiveInviteChan{
			SourceInviteChan: inviteChan.NewWorkerInviteSourceToFannkuchRecursiveSource,
			WorkerInviteChan: inviteChan.NewWorkerInviteNewWorkerToFannkuchRecursiveWorker,
		}

		nestedInviteChannels := &invitations.FannkuchRecursiveNestedInviteChan{
			SourceNestedInviteChan: inviteChan.NewWorkerNestedInviteSourceToFannkuchRecursiveSource,
			WorkerNestedInviteChan: inviteChan.NewWorkerNestedInviteNewWorkerToFannkuchRecursiveWorker,
		}

		FannkuchRecursiveSendCommChannels(wg, inviteChannels, nestedInviteChannels)

		nextWorkerChan := <-inviteChan.NewWorkerInviteNewWorkerToFannkuchRecursiveWorker
		nextInviteChan := <-inviteChan.NewWorkerNestedInviteNewWorkerToFannkuchRecursiveWorker
		workerEnv := env.ToFannkuchRecursiveWorkerEnv()

		workerResult := FannkuchRecursiveWorker(wg, &nextWorkerChan, &nextInviteChan, workerEnv)
		env.ReceiveResultFromWorker(&workerResult)

		result := env.ResultToSource()
		newWorkerChan.SourceResult1 <- result
		env.ResultSentToSource()
	case callbacks.ResultToSource:
		result := env.ResultToSource2()
		newWorkerChan.SourceResult2 <- result
		env.ResultSentToSource2()
	}
	env.Done()
}
