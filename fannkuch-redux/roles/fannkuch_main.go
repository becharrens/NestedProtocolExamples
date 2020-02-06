package roles

import (
	"CodeGenTest/fannkuch-redux/callbacks"
	"CodeGenTest/fannkuch-redux/channels/fannkuch"
	"CodeGenTest/fannkuch-redux/invitations"
	"sync"
)

func StartFannkuchMain(wg *sync.WaitGroup, mainChan *fannkuch.MainChan,
	inviteChan *invitations.FannkuchRoleInviteChan, env callbacks.FannkuchMainEnv) {
	defer wg.Done()
	FannkuchMain(wg, mainChan, inviteChan, env)
}

func FannkuchMain(wg *sync.WaitGroup, mainChan *fannkuch.MainChan, inviteChan *invitations.FannkuchRoleInviteChan,
	env callbacks.FannkuchMainEnv) {
	env.Init()
	task := env.TaskToWorker()

	mainChan.WorkerTask <- task

	env.SentTaskToWorker()

	select {
	case sourceChan := <-inviteChan.WorkerInviteMainToFannkuchRecursiveSource:
		nextInviteChan := <-inviteChan.WorkerNestedInviteMainToFannkuchRecursiveSource
		sourceEnv := env.ToFannkuchRecursiveSourceState()

		sourceResult := FannkuchRecursiveSource(wg, &sourceChan, &nextInviteChan, sourceEnv)
		env.ReceiveResultFromSource(&sourceResult)

		res := <-mainChan.WorkerResult1
		env.ResultFromWorker(&res)
	case res := <-mainChan.WorkerResult2:
		env.ResultFromWorker2(&res)
	}
	env.Done()
}
