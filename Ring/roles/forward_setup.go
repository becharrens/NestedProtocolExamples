package roles

import (
	forward2 "CodeGenTest/Ring/channels/forward"
	"CodeGenTest/Ring/invitations"
	"CodeGenTest/Ring/messages"
	"fmt"
	"sync"
)

func SendCommChannels(wg *sync.WaitGroup, inviteChan *invitations.ForwardInviteChan,
	nestedInviteChan *invitations.ForwardNestedInviteChan) {
	fmt.Println("Setting up nested protocol Forward...")
	sRingNodeMsg := make(chan messages.Msg)
	ringNodeEMsg := make(chan messages.Msg)

	sChan := forward2.SChan{RingNodeMsg: sRingNodeMsg}
	eChan := forward2.EChan{RingNodeMsg: ringNodeEMsg}
	ringNodeChan := forward2.RingNodeChan{
		SMsg: sRingNodeMsg,
		EMsg: ringNodeEMsg,
	}

	roleInviteChan := invitations.ForwardRoleInviteChan{
		RingNodeInviteRingNodeToForwardS:       make(chan forward2.SChan),
		RingNodeInviteEToForwardE:              make(chan forward2.EChan),
		RingNodeNestedInviteEToForwardE:        make(chan invitations.ForwardRoleInviteChan),
		RingNodeNestedInviteRingNodeToForwardS: make(chan invitations.ForwardRoleInviteChan),
	}

	go func() {
		fmt.Println("Inviting role S")
		inviteChan.SInviteChan <- sChan
		nestedInviteChan.SNestedInviteChan <- roleInviteChan
	}()

	go func() {
		fmt.Println("Inviting role E")
		inviteChan.EInviteChan <- eChan
		nestedInviteChan.ENestedInviteChan <- roleInviteChan
	}()

	fmt.Println("Creating role RingNode")
	wg.Add(1)
	go ForwardRingNode(wg, &ringNodeChan, &roleInviteChan)
}
