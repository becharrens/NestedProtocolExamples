package protocol

import (
	"CodeGenTest/ClientServerMiddleware/invitations"
	"CodeGenTest/ClientServerMiddleware/messages"
	"CodeGenTest/ClientServerMiddleware/roles/cms"
	"CodeGenTest/ClientServerMiddleware/roles/contact"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func CMS() {
	// Create random seed
	rand.Seed(time.Now().Unix())

	fmt.Println("Starting protocol CMS...")

	var wg sync.WaitGroup

	reqChan := make(chan messages.Request)
	respChan := make(chan messages.Resp)

	clientChan := cms.ClientChan{
		Req:  reqChan,
		Resp: respChan,
	}
	middleChan := cms.MiddleChan{
		Req:  reqChan,
		Resp: respChan,
	}

	inviteChans := invitations.CommChannels{MiddleToAgent: make(chan contact.AgentChan)}

	wg.Add(2)
	go cms.Client(&wg, &clientChan)
	go cms.Middle(&wg, &middleChan, &inviteChans)
	wg.Wait()

	fmt.Println("CMS protocol Finished")
}
