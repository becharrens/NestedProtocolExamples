package invitations

import "CodeGenTest/ClientServerMiddleware/roles/contact"

type CommChannels struct {
	MiddleToAgent chan contact.AgentChan
}
