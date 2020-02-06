package roles

import (
	"CodeGenTest/Ring/channels/forward"
	"CodeGenTest/Ring/invitations"
	"CodeGenTest/Ring/messages"
	"fmt"
	"sync"
)

func ForwardS(wg *sync.WaitGroup, ringNodeChan *forward.SChan,
	roleInviteChan *invitations.ForwardRoleInviteChan, msg messages.Msg) {
	fmt.Printf("s: forwarding message to RingNode. Current value: %d\n", msg.Val)
	ringNodeChan.RingNodeMsg <- msg
}
