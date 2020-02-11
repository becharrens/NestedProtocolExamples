package protocol

import (
	"CodeGenTest/fannkuch-redux/callbacks"
	"CodeGenTest/fannkuch-redux/channels/fannkuch"
	"CodeGenTest/fannkuch-redux/channels/fannkuchrecursive"
	"CodeGenTest/fannkuch-redux/invitations"
	"CodeGenTest/fannkuch-redux/messages"
	"CodeGenTest/fannkuch-redux/roles"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 3

func Fannkuch() {
	rand.Seed(time.Now().Unix())

	fmt.Print("Starting Fannkuch protocol...\n\n")

	mainWorkerTask := make(chan messages.Task)
	mainWorkerResult1 := make(chan messages.Result)
	mainWorkerResult2 := make(chan messages.Result)

	mainChan := fannkuch.MainChan{
		WorkerTask:    mainWorkerTask,
		WorkerResult1: mainWorkerResult1,
	}
	workerChan := fannkuch.WorkerChan{
		MainTask:    mainWorkerTask,
		MainResult1: mainWorkerResult1,
		MainResult2: mainWorkerResult2,
	}

	inviteChan := invitations.FannkuchRoleInviteChan{
		WorkerInviteMainToFannkuchRecursiveSource:         make(chan fannkuchrecursive.SourceChan),
		WorkerInviteWorkerToFannkuchRecursiveWorker:       make(chan fannkuchrecursive.WorkerChan),
		WorkerNestedInviteMainToFannkuchRecursiveSource:   make(chan invitations.FannkuchRecursiveRoleInviteChan),
		WorkerNestedInviteWorkerToFannkuchRecursiveWorker: make(chan invitations.FannkuchRecursiveRoleInviteChan),
	}

	mainEnv := callbacks.NewFannkuchMainEnv(N)
	workerEnv := callbacks.NewFannkuchWorkerEnv()

	var wg sync.WaitGroup
	wg.Add(2)

	go roles.StartFannkuchMain(&wg, &mainChan, &inviteChan, mainEnv)
	go roles.StartFannkuchWorker(&wg, &workerChan, &inviteChan, workerEnv)

	wg.Wait()
	fmt.Println("\nFannkuch Protocol Fininished...")
}
