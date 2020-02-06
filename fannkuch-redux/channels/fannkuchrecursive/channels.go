package fannkuchrecursive

import "CodeGenTest/fannkuch-redux/messages"

type SourceChan struct {
	NewWorkerResult1 chan messages.Result
	NewWorkerResult2 chan messages.Result
}

type WorkerChan struct {
	NewWorkerTask chan messages.Task
}

type NewWorkerChan struct {
	SourceResult1 chan messages.Result
	SourceResult2 chan messages.Result
	WorkerTask    chan messages.Task
}
