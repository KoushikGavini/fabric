// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/hyperledger/fabric-protos-go-apiv2/common"
	"github.com/hyperledger/fabric-protos-go-apiv2/orderer"
	"google.golang.org/grpc/metadata"
)

type ABBClient struct {
	CloseSendStub        func() error
	closeSendMutex       sync.RWMutex
	closeSendArgsForCall []struct {
	}
	closeSendReturns struct {
		result1 error
	}
	closeSendReturnsOnCall map[int]struct {
		result1 error
	}
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct {
	}
	contextReturns struct {
		result1 context.Context
	}
	contextReturnsOnCall map[int]struct {
		result1 context.Context
	}
	HeaderStub        func() (metadata.MD, error)
	headerMutex       sync.RWMutex
	headerArgsForCall []struct {
	}
	headerReturns struct {
		result1 metadata.MD
		result2 error
	}
	headerReturnsOnCall map[int]struct {
		result1 metadata.MD
		result2 error
	}
	RecvStub        func() (*orderer.BroadcastResponse, error)
	recvMutex       sync.RWMutex
	recvArgsForCall []struct {
	}
	recvReturns struct {
		result1 *orderer.BroadcastResponse
		result2 error
	}
	recvReturnsOnCall map[int]struct {
		result1 *orderer.BroadcastResponse
		result2 error
	}
	RecvMsgStub        func(interface{}) error
	recvMsgMutex       sync.RWMutex
	recvMsgArgsForCall []struct {
		arg1 interface{}
	}
	recvMsgReturns struct {
		result1 error
	}
	recvMsgReturnsOnCall map[int]struct {
		result1 error
	}
	SendStub        func(*common.Envelope) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 *common.Envelope
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	SendMsgStub        func(interface{}) error
	sendMsgMutex       sync.RWMutex
	sendMsgArgsForCall []struct {
		arg1 interface{}
	}
	sendMsgReturns struct {
		result1 error
	}
	sendMsgReturnsOnCall map[int]struct {
		result1 error
	}
	TrailerStub        func() metadata.MD
	trailerMutex       sync.RWMutex
	trailerArgsForCall []struct {
	}
	trailerReturns struct {
		result1 metadata.MD
	}
	trailerReturnsOnCall map[int]struct {
		result1 metadata.MD
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ABBClient) CloseSend() error {
	fake.closeSendMutex.Lock()
	ret, specificReturn := fake.closeSendReturnsOnCall[len(fake.closeSendArgsForCall)]
	fake.closeSendArgsForCall = append(fake.closeSendArgsForCall, struct {
	}{})
	stub := fake.CloseSendStub
	fakeReturns := fake.closeSendReturns
	fake.recordInvocation("CloseSend", []interface{}{})
	fake.closeSendMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ABBClient) CloseSendCallCount() int {
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	return len(fake.closeSendArgsForCall)
}

func (fake *ABBClient) CloseSendCalls(stub func() error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = stub
}

func (fake *ABBClient) CloseSendReturns(result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	fake.closeSendReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) CloseSendReturnsOnCall(i int, result1 error) {
	fake.closeSendMutex.Lock()
	defer fake.closeSendMutex.Unlock()
	fake.CloseSendStub = nil
	if fake.closeSendReturnsOnCall == nil {
		fake.closeSendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeSendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) Context() context.Context {
	fake.contextMutex.Lock()
	ret, specificReturn := fake.contextReturnsOnCall[len(fake.contextArgsForCall)]
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct {
	}{})
	stub := fake.ContextStub
	fakeReturns := fake.contextReturns
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ABBClient) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *ABBClient) ContextCalls(stub func() context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = stub
}

func (fake *ABBClient) ContextReturns(result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *ABBClient) ContextReturnsOnCall(i int, result1 context.Context) {
	fake.contextMutex.Lock()
	defer fake.contextMutex.Unlock()
	fake.ContextStub = nil
	if fake.contextReturnsOnCall == nil {
		fake.contextReturnsOnCall = make(map[int]struct {
			result1 context.Context
		})
	}
	fake.contextReturnsOnCall[i] = struct {
		result1 context.Context
	}{result1}
}

func (fake *ABBClient) Header() (metadata.MD, error) {
	fake.headerMutex.Lock()
	ret, specificReturn := fake.headerReturnsOnCall[len(fake.headerArgsForCall)]
	fake.headerArgsForCall = append(fake.headerArgsForCall, struct {
	}{})
	stub := fake.HeaderStub
	fakeReturns := fake.headerReturns
	fake.recordInvocation("Header", []interface{}{})
	fake.headerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ABBClient) HeaderCallCount() int {
	fake.headerMutex.RLock()
	defer fake.headerMutex.RUnlock()
	return len(fake.headerArgsForCall)
}

func (fake *ABBClient) HeaderCalls(stub func() (metadata.MD, error)) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = stub
}

func (fake *ABBClient) HeaderReturns(result1 metadata.MD, result2 error) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = nil
	fake.headerReturns = struct {
		result1 metadata.MD
		result2 error
	}{result1, result2}
}

func (fake *ABBClient) HeaderReturnsOnCall(i int, result1 metadata.MD, result2 error) {
	fake.headerMutex.Lock()
	defer fake.headerMutex.Unlock()
	fake.HeaderStub = nil
	if fake.headerReturnsOnCall == nil {
		fake.headerReturnsOnCall = make(map[int]struct {
			result1 metadata.MD
			result2 error
		})
	}
	fake.headerReturnsOnCall[i] = struct {
		result1 metadata.MD
		result2 error
	}{result1, result2}
}

func (fake *ABBClient) Recv() (*orderer.BroadcastResponse, error) {
	fake.recvMutex.Lock()
	ret, specificReturn := fake.recvReturnsOnCall[len(fake.recvArgsForCall)]
	fake.recvArgsForCall = append(fake.recvArgsForCall, struct {
	}{})
	stub := fake.RecvStub
	fakeReturns := fake.recvReturns
	fake.recordInvocation("Recv", []interface{}{})
	fake.recvMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ABBClient) RecvCallCount() int {
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	return len(fake.recvArgsForCall)
}

func (fake *ABBClient) RecvCalls(stub func() (*orderer.BroadcastResponse, error)) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = stub
}

func (fake *ABBClient) RecvReturns(result1 *orderer.BroadcastResponse, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	fake.recvReturns = struct {
		result1 *orderer.BroadcastResponse
		result2 error
	}{result1, result2}
}

func (fake *ABBClient) RecvReturnsOnCall(i int, result1 *orderer.BroadcastResponse, result2 error) {
	fake.recvMutex.Lock()
	defer fake.recvMutex.Unlock()
	fake.RecvStub = nil
	if fake.recvReturnsOnCall == nil {
		fake.recvReturnsOnCall = make(map[int]struct {
			result1 *orderer.BroadcastResponse
			result2 error
		})
	}
	fake.recvReturnsOnCall[i] = struct {
		result1 *orderer.BroadcastResponse
		result2 error
	}{result1, result2}
}

func (fake *ABBClient) RecvMsg(arg1 interface{}) error {
	fake.recvMsgMutex.Lock()
	ret, specificReturn := fake.recvMsgReturnsOnCall[len(fake.recvMsgArgsForCall)]
	fake.recvMsgArgsForCall = append(fake.recvMsgArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	stub := fake.RecvMsgStub
	fakeReturns := fake.recvMsgReturns
	fake.recordInvocation("RecvMsg", []interface{}{arg1})
	fake.recvMsgMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ABBClient) RecvMsgCallCount() int {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	return len(fake.recvMsgArgsForCall)
}

func (fake *ABBClient) RecvMsgCalls(stub func(interface{}) error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = stub
}

func (fake *ABBClient) RecvMsgArgsForCall(i int) interface{} {
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	argsForCall := fake.recvMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ABBClient) RecvMsgReturns(result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	fake.recvMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) RecvMsgReturnsOnCall(i int, result1 error) {
	fake.recvMsgMutex.Lock()
	defer fake.recvMsgMutex.Unlock()
	fake.RecvMsgStub = nil
	if fake.recvMsgReturnsOnCall == nil {
		fake.recvMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.recvMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) Send(arg1 *common.Envelope) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 *common.Envelope
	}{arg1})
	stub := fake.SendStub
	fakeReturns := fake.sendReturns
	fake.recordInvocation("Send", []interface{}{arg1})
	fake.sendMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ABBClient) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *ABBClient) SendCalls(stub func(*common.Envelope) error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = stub
}

func (fake *ABBClient) SendArgsForCall(i int) *common.Envelope {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	argsForCall := fake.sendArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ABBClient) SendReturns(result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) SendReturnsOnCall(i int, result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) SendMsg(arg1 interface{}) error {
	fake.sendMsgMutex.Lock()
	ret, specificReturn := fake.sendMsgReturnsOnCall[len(fake.sendMsgArgsForCall)]
	fake.sendMsgArgsForCall = append(fake.sendMsgArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	stub := fake.SendMsgStub
	fakeReturns := fake.sendMsgReturns
	fake.recordInvocation("SendMsg", []interface{}{arg1})
	fake.sendMsgMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ABBClient) SendMsgCallCount() int {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	return len(fake.sendMsgArgsForCall)
}

func (fake *ABBClient) SendMsgCalls(stub func(interface{}) error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = stub
}

func (fake *ABBClient) SendMsgArgsForCall(i int) interface{} {
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	argsForCall := fake.sendMsgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ABBClient) SendMsgReturns(result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	fake.sendMsgReturns = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) SendMsgReturnsOnCall(i int, result1 error) {
	fake.sendMsgMutex.Lock()
	defer fake.sendMsgMutex.Unlock()
	fake.SendMsgStub = nil
	if fake.sendMsgReturnsOnCall == nil {
		fake.sendMsgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMsgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ABBClient) Trailer() metadata.MD {
	fake.trailerMutex.Lock()
	ret, specificReturn := fake.trailerReturnsOnCall[len(fake.trailerArgsForCall)]
	fake.trailerArgsForCall = append(fake.trailerArgsForCall, struct {
	}{})
	stub := fake.TrailerStub
	fakeReturns := fake.trailerReturns
	fake.recordInvocation("Trailer", []interface{}{})
	fake.trailerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ABBClient) TrailerCallCount() int {
	fake.trailerMutex.RLock()
	defer fake.trailerMutex.RUnlock()
	return len(fake.trailerArgsForCall)
}

func (fake *ABBClient) TrailerCalls(stub func() metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = stub
}

func (fake *ABBClient) TrailerReturns(result1 metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = nil
	fake.trailerReturns = struct {
		result1 metadata.MD
	}{result1}
}

func (fake *ABBClient) TrailerReturnsOnCall(i int, result1 metadata.MD) {
	fake.trailerMutex.Lock()
	defer fake.trailerMutex.Unlock()
	fake.TrailerStub = nil
	if fake.trailerReturnsOnCall == nil {
		fake.trailerReturnsOnCall = make(map[int]struct {
			result1 metadata.MD
		})
	}
	fake.trailerReturnsOnCall[i] = struct {
		result1 metadata.MD
	}{result1}
}

func (fake *ABBClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeSendMutex.RLock()
	defer fake.closeSendMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	fake.headerMutex.RLock()
	defer fake.headerMutex.RUnlock()
	fake.recvMutex.RLock()
	defer fake.recvMutex.RUnlock()
	fake.recvMsgMutex.RLock()
	defer fake.recvMsgMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	fake.sendMsgMutex.RLock()
	defer fake.sendMsgMutex.RUnlock()
	fake.trailerMutex.RLock()
	defer fake.trailerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ABBClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
