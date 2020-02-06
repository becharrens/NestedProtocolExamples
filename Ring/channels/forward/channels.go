package forward

import "CodeGenTest/Ring/messages"

type EChan struct {
    RingNodeMsg chan messages.Msg
}

type SChan struct {
    RingNodeMsg chan messages.Msg
}

type RingNodeChan struct {
    SMsg chan messages.Msg
    EMsg chan messages.Msg
}