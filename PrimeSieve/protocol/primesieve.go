package protocol

import (
	"CodeGenTest/PrimeSieve/callbacks"
	primesieve2 "CodeGenTest/PrimeSieve/channels/primesieve"
	"CodeGenTest/PrimeSieve/channels/sieve"
	"CodeGenTest/PrimeSieve/invitations"
	"CodeGenTest/PrimeSieve/messages/primesieve"
	"CodeGenTest/PrimeSieve/roles"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func PrimeSieve() {
	rand.Seed(time.Now().Unix())

	n := rand.Intn(100) + 2

	fmt.Print("Starting Prime Sieve protocol...\n\n")

	masterWorkerFirstPrime := make(chan primesieve.FirstPrime)
	masterWorkerUbound := make(chan primesieve.UBound)
	masterWorkerPrime := make(chan primesieve.Prime)
	masterWorkerFinish := make(chan primesieve.Finish)

	masterChan := primesieve2.MasterChan{
		WorkerFirstPrime: masterWorkerFirstPrime,
		WorkerUBound:     masterWorkerUbound,
		WorkerPrime:      masterWorkerPrime,
		WorkerFinish:     masterWorkerFinish,
	}

	workerChan := primesieve2.WorkerChan{
		MasterFirstPrime: masterWorkerFirstPrime,
		MasterUBound:     masterWorkerUbound,
		MasterPrime:      masterWorkerPrime,
		MasterFinish:     masterWorkerFinish,
	}

	workerInviteMasterToSieveM := make(chan sieve.MChan)
	workerInviteWorkerToSieveW1 := make(chan sieve.W1Chan)
	workerInviteMasterToSieveMInviteChan := make(chan invitations.SieveMInviteChan)
	workerInviteWorkerToSieveW1InviteChan := make(chan invitations.SieveW1InviteChan)

	masterInviteChan := invitations.PrimeSieveMasterInviteChan{
		WorkerInviteToSieveM:           workerInviteMasterToSieveM,
		WorkerInviteToSieveMInviteChan: workerInviteMasterToSieveMInviteChan,
	}

	workerInviteChan := invitations.PrimeSieveWorkerInviteChan{
		InviteWorkerToSieveW1:           workerInviteWorkerToSieveW1,
		InviteWorkerToSieveW1InviteChan: workerInviteWorkerToSieveW1InviteChan,
		InviteMasterToSieveM:            workerInviteMasterToSieveM,
		InviteMasterToSieveMInviteChan:  workerInviteMasterToSieveMInviteChan,
	}

	masterEnv := callbacks.NewPrimeSieveMasterEnv(n)
	workerEnv := callbacks.NewPrimeSieveWorkerEnv()

	var wg sync.WaitGroup
	wg.Add(2)

	go roles.StartPrimeSieveMaster(&wg, &masterChan, &masterInviteChan, masterEnv)
	go roles.StartPrimeSieveWorker(&wg, &workerChan, &workerInviteChan, workerEnv)

	wg.Wait()
	fmt.Println("\nPrime Sieve Protocol Fininished...")
}
