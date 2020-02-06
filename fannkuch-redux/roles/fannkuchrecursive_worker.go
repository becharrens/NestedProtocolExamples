package roles

import (
	"CodeGenTest/fannkuch-redux/callbacks"
	fannkuchrecursive2 "CodeGenTest/fannkuch-redux/callbacks/result/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/channels/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/invitations"
	"sync"
)

func FannkuchRecursiveWorker(wg *sync.WaitGroup,
	workerChan *fannkuchrecursive.WorkerChan,
	inviteChan *invitations.FannkuchRecursiveRoleInviteChan,
	env callbacks.FannkuchRecursiveWorkerEnv) fannkuchrecursive2.WorkerResult {
	env.Init()

	task := env.TaskToNewWorker()
	workerChan.NewWorkerTask <- task
	env.SentTaskToNewWorker()

	return env.Done()
}
