package protocol

import (
	forward2 "CodeGenTest/Ring/channels/forward"
	"CodeGenTest/Ring/channels/ring"
	"CodeGenTest/Ring/invitations"
	"CodeGenTest/Ring/messages"
	"CodeGenTest/Ring/roles"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Ring() {
	rand.Seed(time.Now().Unix())

	fmt.Print("Starting Ring protocol...\n\n")

	startEndMsg := make(chan messages.Msg)
	startChan := ring.StartChan{EndMsg: startEndMsg}
	endChan := ring.EndChan{StartMsg: startEndMsg}

	roleInviteChan := invitations.RingRoleInviteChan{
		StartInviteStartToS:       make(chan forward2.SChan),
		StartInviteEndToE:         make(chan forward2.EChan),
		StartNestedInviteStartToS: make(chan invitations.ForwardRoleInviteChan),
		StartNestedInviteEndToE:   make(chan invitations.ForwardRoleInviteChan),
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go roles.StartRingStart(&wg, &startChan, &roleInviteChan)
	go roles.StartRingEnd(&wg, &endChan, &roleInviteChan)

	wg.Wait()
	fmt.Println("\nRing Protocol Fininished...")
}
