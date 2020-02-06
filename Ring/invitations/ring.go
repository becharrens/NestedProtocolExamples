package invitations

import (
	"CodeGenTest/Ring/channels/forward"
)

type RingRoleInviteChan struct {
	StartInviteStartToS       chan forward.SChan
	StartInviteEndToE         chan forward.EChan
	StartNestedInviteStartToS chan ForwardRoleInviteChan
	StartNestedInviteEndToE   chan ForwardRoleInviteChan
}
