package primesieve

import "CodeGenTest/PrimeSieve/messages/primesieve"

type MasterChan struct {
	WorkerFirstPrime chan primesieve.FirstPrime
	WorkerUBound     chan primesieve.UBound
	WorkerPrime      chan primesieve.Prime
	WorkerFinish     chan primesieve.Finish
}

type WorkerChan struct {
	MasterFirstPrime chan primesieve.FirstPrime
	MasterUBound     chan primesieve.UBound
	MasterPrime      chan primesieve.Prime
	MasterFinish     chan primesieve.Finish
}
