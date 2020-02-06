package roles

import (
	"CodeGenTest/Ring/channels/forward"
	"CodeGenTest/Ring/invitations"
	"CodeGenTest/Ring/messages"
	"fmt"
	"sync"
)

func ForwardRingNode(wg *sync.WaitGroup, ringNodeChan *forward.RingNodeChan,
	roleInviteChan *invitations.ForwardRoleInviteChan) {
	defer wg.Done()
	msg := <-ringNodeChan.SMsg
	fmt.Printf("ringNode: received message from S. Current value: %d\n", msg.Val)
	if msg.Val > 2 {
		inviteChan := &invitations.ForwardInviteChan{
			EInviteChan: roleInviteChan.RingNodeInviteEToForwardE,
			SInviteChan: roleInviteChan.RingNodeInviteRingNodeToForwardS,
		}
		nestedInviteChan := &invitations.ForwardNestedInviteChan{
			ENestedInviteChan: roleInviteChan.RingNodeNestedInviteEToForwardE,
			SNestedInviteChan: roleInviteChan.RingNodeNestedInviteRingNodeToForwardS,
		}
		fmt.Println("ringNode: forwarding msg to next node in the ring")
		SendCommChannels(wg, inviteChan, nestedInviteChan)

		nextSInvite := <-roleInviteChan.RingNodeInviteRingNodeToForwardS
		nextRoleInviteChan := <-roleInviteChan.RingNodeNestedInviteRingNodeToForwardS

		fmt.Println("ringNode: accept role S in Forward")
		ForwardS(wg, &nextSInvite, &nextRoleInviteChan, messages.Msg{Val: msg.Val - 1})
	} else {
		fmt.Println("ringNode: sending last message to e")
		ringNodeChan.EMsg <- messages.Msg{Val: msg.Val - 1}
	}
}
