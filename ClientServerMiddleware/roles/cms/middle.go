package cms

import (
	"CodeGenTest/ClientServerMiddleware/invitations"
	"CodeGenTest/ClientServerMiddleware/messages"
	"CodeGenTest/ClientServerMiddleware/roles/contact"
	"fmt"
	"math/rand"
	"sync"
)

type MiddleChan struct {
	Req  chan messages.Request
	Resp chan messages.Resp
}

func Middle(group *sync.WaitGroup, middleChan *MiddleChan, channels *invitations.CommChannels) {
	defer group.Done()

	req := <-middleChan.Req
	fmt.Printf("Middle: received reqest '%s'\n", req.Req)

	if rand.Intn(2)%2 == 0 {
		// TODO: Invitations
		//  - send channels for existing participants
		//  - External channels are implicit (currently). Might be needed in
		// a distributed setting
		go contact.SendCommChannels(channels.MiddleToAgent)
		agentChan := <-channels.MiddleToAgent

		// TODO: How to implement return values from nested protocols:
		// - Depend on callbacks and push responsability onto the programmer
		// - Allow user to specify caller of protocol as the participant who
		//   always returns the value

		fmt.Println("Middle: Playing Agent role in Contact protocol")
		resp := contact.ContactAgent(&agentChan, req)
		middleChan.Resp <- resp
		fmt.Printf("Middle: sent response '%s'\n", resp.Response)
	} else {
		// TODO: input
		resp := messages.Resp{Response: "Cached"}
		middleChan.Resp <- resp
		fmt.Printf("Middle: sent response '%s'\n", resp.Response)
	}
}
