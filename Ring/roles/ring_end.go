package roles

import (
	"CodeGenTest/Ring/channels/ring"
	"CodeGenTest/Ring/invitations"
	"CodeGenTest/Ring/messages"
	"fmt"
	"sync"
)

func StartRingEnd(wg *sync.WaitGroup, endChan *ring.EndChan, inviteChan *invitations.RingRoleInviteChan) {
	defer wg.Done()
	RingEnd(wg, inviteChan, endChan)
}

func RingEnd(wg *sync.WaitGroup, inviteChan *invitations.RingRoleInviteChan, endChan *ring.EndChan) {
	var msg messages.Msg
	select {
	case nextEChan := <-inviteChan.StartInviteEndToE:
		nextInviteChan := <-inviteChan.StartNestedInviteEndToE
		fmt.Println("end: accept role E in Forward")
		msg = ForwardE(wg, &nextEChan, &nextInviteChan)
		fmt.Print("end: received message from ring. ")
	case msg = <-endChan.StartMsg:
		fmt.Print("end: received message from start. ")
	}
	fmt.Printf("Current value: %d\n", msg.Val)
	fmt.Println("end: forwarding message back to start")
	endChan.StartMsg <- messages.Msg{Val: msg.Val - 1}
}
