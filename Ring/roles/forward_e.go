package roles

import (
	"CodeGenTest/Ring/channels/forward"
	"CodeGenTest/Ring/invitations"
	"CodeGenTest/Ring/messages"
	"fmt"
	"sync"
)

func ForwardE(wg *sync.WaitGroup, eChan *forward.EChan, inviteChan *invitations.ForwardRoleInviteChan) messages.Msg {
	select {
	case msg := <-eChan.RingNodeMsg:
		fmt.Printf("e: Received message from RingNode. Current value: %d\n", msg.Val)
		return msg
	case nextEChan := <-inviteChan.RingNodeInviteEToForwardE:
		fmt.Println("e: Accept invite for role E in Forward")
		nextEInviteChan := <-inviteChan.RingNodeNestedInviteEToForwardE
		return ForwardE(wg, &nextEChan, &nextEInviteChan)
	}
}
