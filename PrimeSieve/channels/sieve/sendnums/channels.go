package sendnums

import "CodeGenTest/PrimeSieve/messages/sieve/sendnums"

type SChan struct {
	RNum chan sendnums.Num
	REnd chan sendnums.End
}

type RChan struct {
	SNum chan sendnums.Num
	SEnd chan sendnums.End
}
