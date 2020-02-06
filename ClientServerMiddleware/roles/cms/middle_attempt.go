package cms

//
// import (
// 	"CodeGenTest/ClientServerMiddleware/messages"
// 	"CodeGenTest/ClientServerMiddleware/protocol"
// 	"sync"
// )
//
// // type MiddleChan struct {
// // 	Req  chan messages.Request
// // 	Resp chan messages.Resp
// // }
//
// func StartState() *MS1 {
// 	return &MS1{}
// }
//
// type MS1 struct {
// 	Req messages.Request
// }
//
// func (s *MS1) Callback() {
// 	panic("implement me")
// }
//
// type MSChoice0 struct {
// 	Branch int
// }
//
// type MS2B0 struct {
// 	Resp messages.Resp
// }
//
// func (s *MS2B0) Callback() {
// 	panic("implement me")
// }
//
// type MS2B1 struct {
// }
//
// type MS3 struct {
// }
//
// type MSEnd struct {
// }
//
// func (end *MSEnd) Callback() {
// 	panic("implement me")
// }
//
// func NewMSChoice0(ms1 *MS1) *MSChoice0 {
// 	return &MSChoice0{Branch: 1}
// }
//
// func NewMS2B1(msChoice *MSChoice0) *MS2B1 {
// 	return &MS2B1{}
// }
//
// func NewMS2B0(msChoice *MSChoice0) *MS2B0 {
// 	return &MS2B0{}
// }
//
// func NewMEnd(ms *MS2B0) *MSEnd {
// 	return &MSEnd{}
// }
//
// func Middle1(wg sync.WaitGroup, state *MS1, middleChan *MiddleChan) {
// 	defer wg.Done()
// 	req := <-middleChan.Req
// 	state.Req = req
// 	state.Callback()
// 	next := NewMSChoice0(state)
// 	middleChooseBranch(next, middleChan)
// }
//
// func middleChooseBranch(state *MSChoice0, middleChan *MiddleChan) {
// 	switch state.Branch {
// 	case 0:
// 		next := NewMS2B0(state)
// 		middleBra0(next, middleChan)
// 	case 1:
// 		next := NewMS2B1(state)
// 		middleBra1(next, middleChan)
// 	}
// }
//
// func middleBra0(ms2b0 *MS2B0, middleChan *MiddleChan) {
// 	middleChan.Resp <- ms2b0.Resp
// 	ms2b0.Callback()
//
// }
//
// func middleBra1(ms2b1 *MS2B1, middleChan *MiddleChan) {
// 	res := make(chan messages.Resp)
// 	go protocol.Contact(res)
// 	// wait for it to finish and get result
//
// 	ms2b1.Callback()
// 	middleToClient()
// }
//
// func middleEnd(end *MSEnd, middleChan *MiddleChan) {
// 	end.Callback()
// }
