package roles

import (
	"CodeGenTest/fannkuch-redux/callbacks"
	fannkuchrecursive2 "CodeGenTest/fannkuch-redux/callbacks/result/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/channels/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/invitations"
	"fmt"
	"sync"
)

func FannkuchRecursiveSource(wg *sync.WaitGroup,
	sourceChan *fannkuchrecursive.SourceChan,
	inviteChan *invitations.FannkuchRecursiveRoleInviteChan,
	env callbacks.FannkuchRecursiveSourceEnv) fannkuchrecursive2.SourceResult {

	env.Init()

	select {
	case nextSourceChan := <-inviteChan.NewWorkerInviteSourceToFannkuchRecursiveSource:
		fmt.Println("source invited to source.")
		nextInviteChan := <-inviteChan.NewWorkerNestedInviteSourceToFannkuchRecursiveSource
		sourceEnv := env.ToFannkuchRecursiveSourceEnv()
		sourceResult := FannkuchRecursiveSource(wg, &nextSourceChan, &nextInviteChan, sourceEnv)
		env.ReceiveResultFromSource(&sourceResult)

		fmt.Println("source waiting for result")
		result := <-sourceChan.NewWorkerResult1
		fmt.Printf("source received result: %d, %d\n", result.MaxFlips, result.CheckSum)
		env.ResultFromNewWorker(&result)
	case result := <-sourceChan.NewWorkerResult2:
		env.ResultFromNewWorker2(&result)
	}

	return env.Done()
}
