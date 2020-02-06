package contact

import (
	"CodeGenTest/ClientServerMiddleware/messages"
	"fmt"
)

func SendCommChannels(agentInvite chan AgentChan) {
	fmt.Println("Sending channels for Contact protocol")

	reqChan := make(chan messages.Request)
	respChan := make(chan messages.Resp)

	agentChan := AgentChan{
		Req:  reqChan,
		Resp: respChan,
	}
	serverChan := ServerChan{
		Req:  reqChan,
		Resp: respChan,
	}

	go func() { agentInvite <- agentChan }()

	fmt.Println("Starting new server")
	go ContactServer(&serverChan)
}
