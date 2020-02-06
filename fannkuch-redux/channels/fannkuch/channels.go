package fannkuch

import "CodeGenTest/fannkuch-redux/messages"

type MainChan struct {
	WorkerTask    chan messages.Task
	WorkerResult1 chan messages.Result
	WorkerResult2 chan messages.Result
}

type WorkerChan struct {
	MainTask    chan messages.Task
	MainResult1 chan messages.Result
	MainResult2 chan messages.Result
}
