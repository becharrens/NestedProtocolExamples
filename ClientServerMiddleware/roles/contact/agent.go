package contact

import (
	"CodeGenTest/ClientServerMiddleware/messages"
	"fmt"
)

type AgentChan struct {
	Req  chan messages.Request
	Resp chan messages.Resp
}

// func ContactAgent(agentChan *AgentChan, req messages.Request) {
// 	// Subprotocols don't need to be waited on, right?
//
// 	fmt.Printf("Agent sending req '%s' to Server\n", req.Req)
// 	agentChan.Req <- req
// 	resp := <-agentChan.Resp
// 	fmt.Printf("Agent received resp '%s' from Server", resp.Response)
// }

func ContactAgent(agentChan *AgentChan, req messages.Request) messages.Resp {
	// Subprotocols don't need to be waited on, right?

	fmt.Printf("Agent: sending req '%s' to Server\n", req.Req)
	agentChan.Req <- req
	resp := <-agentChan.Resp
	fmt.Printf("Agent: received resp '%s' from Server\n", resp.Response)
	return resp
}
