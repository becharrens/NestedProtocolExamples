package roles

import (
	"CodeGenTest/Ring/channels/ring"
	"CodeGenTest/Ring/invitations"
	"CodeGenTest/Ring/messages"
	"fmt"
	"math/rand"
	"sync"
)

func StartRingStart(wg *sync.WaitGroup, startChan *ring.StartChan, roleInviteChan *invitations.RingRoleInviteChan) {
	defer wg.Done()
	RingStart(wg, startChan, roleInviteChan)
}

func RingStart(wg *sync.WaitGroup, startChan *ring.StartChan, roleInviteChan *invitations.RingRoleInviteChan) {
	ringSize := rand.Intn(9) + 2

	fmt.Printf("start: creating ring of size: %d\n", ringSize)
	fmt.Printf("start: initial value = %d\n", ringSize)
	msg := messages.Msg{Val: ringSize - 1}
	if ringSize == 2 {
		fmt.Println("start: send message to the other node in the ring")
		startChan.EndMsg <- msg
	} else {
		inviteChan := &invitations.ForwardInviteChan{
			EInviteChan: roleInviteChan.StartInviteEndToE,
			SInviteChan: roleInviteChan.StartInviteStartToS,
		}
		nestedInviteChan := &invitations.ForwardNestedInviteChan{
			ENestedInviteChan: roleInviteChan.StartNestedInviteEndToE,
			SNestedInviteChan: roleInviteChan.StartNestedInviteStartToS,
		}
		fmt.Println("start: ring size is greater than two. Creating ring with dynamic participants")
		SendCommChannels(wg, inviteChan, nestedInviteChan)

		nextSInvite := <-roleInviteChan.StartInviteStartToS
		nextRoleInviteChan := <-roleInviteChan.StartNestedInviteStartToS
		fmt.Println("start: accept role S in Forward")
		ForwardS(wg, &nextSInvite, &nextRoleInviteChan, messages.Msg{Val: msg.Val})
	}
	endMsg := <-startChan.EndMsg
	fmt.Printf("start: value returned by the ring is %d\n", endMsg.Val)
}
