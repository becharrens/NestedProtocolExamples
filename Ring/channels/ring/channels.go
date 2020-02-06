package ring

import "CodeGenTest/Ring/messages"

type StartChan struct {
    EndMsg chan messages.Msg
}

type EndChan struct {
    StartMsg chan messages.Msg
}
