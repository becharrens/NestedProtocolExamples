package cms

import (
	"CodeGenTest/ClientServerMiddleware/messages"
	"fmt"
	"sync"
)

type ClientChan struct {
	Req  chan messages.Request
	Resp chan messages.Resp
}

func Client(wg *sync.WaitGroup, clientChan *ClientChan) {
	// TODO: read input?
	defer wg.Done()
	request := messages.Request{Req: "Req0"}
	fmt.Printf("Client: sending req '%s'\n", request.Req)
	clientChan.Req <- request
	resp := <-clientChan.Resp
	fmt.Printf("Client: received response '%s'\n", resp.Response)
}
