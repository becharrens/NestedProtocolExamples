package invitations

import "CodeGenTest/Ring/channels/forward"

// struct used in setup to identify who to send invitations to
type ForwardInviteChan struct {
	EInviteChan chan forward.EChan
	SInviteChan chan forward.SChan
}

// struct used in setup to identify who to send nested invitation channels to
type ForwardNestedInviteChan struct {
	ENestedInviteChan chan ForwardRoleInviteChan
	SNestedInviteChan chan ForwardRoleInviteChan
}

type ForwardRoleInviteChan struct {
	RingNodeInviteRingNodeToForwardS       chan forward.SChan
	RingNodeInviteEToForwardE              chan forward.EChan
	RingNodeNestedInviteEToForwardE        chan ForwardRoleInviteChan
	RingNodeNestedInviteRingNodeToForwardS chan ForwardRoleInviteChan
}
