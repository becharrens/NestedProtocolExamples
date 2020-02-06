package contact

import (
	"CodeGenTest/ClientServerMiddleware/messages"
	"fmt"
)

type ServerChan struct {
	Req  chan messages.Request
	Resp chan messages.Resp
}

func ContactServer(serverChan *ServerChan) {
	// How would external channels work in a distributed system
	// (not goroutines)
	req := <-serverChan.Req
	fmt.Printf("Server: received req '%s'\n", req.Req)

	resp := messages.Resp{Response: "Resp"}
	serverChan.Resp <- resp
	fmt.Printf("Server: sent response '%s'\n", resp.Response)
}
