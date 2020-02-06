package roles

import (
	"CodeGenTest/fannkuch-redux/callbacks"
	"CodeGenTest/fannkuch-redux/channels/fannkuch"
	"CodeGenTest/fannkuch-redux/invitations"
	"fmt"
	"sync"
)

func StartFannkuchWorker(wg *sync.WaitGroup, workerChan *fannkuch.WorkerChan,
	inviteChan *invitations.FannkuchRoleInviteChan, env callbacks.FannkuchWorkerEnv) {
	defer wg.Done()
	FannkuchWorker(wg, workerChan, inviteChan, env)
}

func FannkuchWorker(wg *sync.WaitGroup, workerChan *fannkuch.WorkerChan,
	inviteChan *invitations.FannkuchRoleInviteChan, env callbacks.FannkuchWorkerEnv) {
	env.Init()

	task := <-workerChan.MainTask
	env.TaskFromMain(&task)

	choice := env.WorkerChoice()
	switch choice {
	case callbacks.CallFannkuchRecursive:
		fmt.Println("Worker recurse")

		inviteChannels := &invitations.FannkuchRecursiveInviteChan{
			SourceInviteChan: inviteChan.WorkerInviteMainToFannkuchRecursiveSource,
			WorkerInviteChan: inviteChan.WorkerInviteWorkerToFannkuchRecursiveWorker,
		}

		nestedInviteChannels := &invitations.FannkuchRecursiveNestedInviteChan{
			SourceNestedInviteChan: inviteChan.WorkerNestedInviteMainToFannkuchRecursiveSource,
			WorkerNestedInviteChan: inviteChan.WorkerNestedInviteWorkerToFannkuchRecursiveWorker,
		}
		fmt.Println("Worker invite main to source")
		FannkuchRecursiveSendCommChannels(wg, inviteChannels, nestedInviteChannels)

		nextWorkerChan := <-inviteChan.WorkerInviteWorkerToFannkuchRecursiveWorker
		nextInviteChan := <-inviteChan.WorkerNestedInviteWorkerToFannkuchRecursiveWorker
		workerEnv := env.ToFannkuchRecursiveWorkerState()
		workerResult := FannkuchRecursiveWorker(wg, &nextWorkerChan, &nextInviteChan, workerEnv)

		env.ReceiveResultFromWorker(&workerResult)

		res := env.ResultToMain()
		workerChan.MainResult1 <- res
		env.SentResultToMain()

	case callbacks.ResultToMain:
		res := env.ResultToMain2()
		workerChan.MainResult2 <- res
		env.SentResultSentToMain2()
	}

	env.Done()
}
