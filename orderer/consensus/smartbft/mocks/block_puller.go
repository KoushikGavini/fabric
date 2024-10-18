// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go-apiv2/common"
	"github.com/hyperledger/fabric/orderer/consensus/smartbft"
)

type FakeBlockPuller struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	HeightsByEndpointsStub        func() (map[string]uint64, string, error)
	heightsByEndpointsMutex       sync.RWMutex
	heightsByEndpointsArgsForCall []struct {
	}
	heightsByEndpointsReturns struct {
		result1 map[string]uint64
		result2 string
		result3 error
	}
	heightsByEndpointsReturnsOnCall map[int]struct {
		result1 map[string]uint64
		result2 string
		result3 error
	}
	PullBlockStub        func(uint64) *common.Block
	pullBlockMutex       sync.RWMutex
	pullBlockArgsForCall []struct {
		arg1 uint64
	}
	pullBlockReturns struct {
		result1 *common.Block
	}
	pullBlockReturnsOnCall map[int]struct {
		result1 *common.Block
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockPuller) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeBlockPuller) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeBlockPuller) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeBlockPuller) HeightsByEndpoints() (map[string]uint64, string, error) {
	fake.heightsByEndpointsMutex.Lock()
	ret, specificReturn := fake.heightsByEndpointsReturnsOnCall[len(fake.heightsByEndpointsArgsForCall)]
	fake.heightsByEndpointsArgsForCall = append(fake.heightsByEndpointsArgsForCall, struct {
	}{})
	stub := fake.HeightsByEndpointsStub
	fakeReturns := fake.heightsByEndpointsReturns
	fake.recordInvocation("HeightsByEndpoints", []interface{}{})
	fake.heightsByEndpointsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeBlockPuller) HeightsByEndpointsCallCount() int {
	fake.heightsByEndpointsMutex.RLock()
	defer fake.heightsByEndpointsMutex.RUnlock()
	return len(fake.heightsByEndpointsArgsForCall)
}

func (fake *FakeBlockPuller) HeightsByEndpointsCalls(stub func() (map[string]uint64, string, error)) {
	fake.heightsByEndpointsMutex.Lock()
	defer fake.heightsByEndpointsMutex.Unlock()
	fake.HeightsByEndpointsStub = stub
}

func (fake *FakeBlockPuller) HeightsByEndpointsReturns(result1 map[string]uint64, result2 string, result3 error) {
	fake.heightsByEndpointsMutex.Lock()
	defer fake.heightsByEndpointsMutex.Unlock()
	fake.HeightsByEndpointsStub = nil
	fake.heightsByEndpointsReturns = struct {
		result1 map[string]uint64
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBlockPuller) HeightsByEndpointsReturnsOnCall(i int, result1 map[string]uint64, result2 string, result3 error) {
	fake.heightsByEndpointsMutex.Lock()
	defer fake.heightsByEndpointsMutex.Unlock()
	fake.HeightsByEndpointsStub = nil
	if fake.heightsByEndpointsReturnsOnCall == nil {
		fake.heightsByEndpointsReturnsOnCall = make(map[int]struct {
			result1 map[string]uint64
			result2 string
			result3 error
		})
	}
	fake.heightsByEndpointsReturnsOnCall[i] = struct {
		result1 map[string]uint64
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBlockPuller) PullBlock(arg1 uint64) *common.Block {
	fake.pullBlockMutex.Lock()
	ret, specificReturn := fake.pullBlockReturnsOnCall[len(fake.pullBlockArgsForCall)]
	fake.pullBlockArgsForCall = append(fake.pullBlockArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.PullBlockStub
	fakeReturns := fake.pullBlockReturns
	fake.recordInvocation("PullBlock", []interface{}{arg1})
	fake.pullBlockMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockPuller) PullBlockCallCount() int {
	fake.pullBlockMutex.RLock()
	defer fake.pullBlockMutex.RUnlock()
	return len(fake.pullBlockArgsForCall)
}

func (fake *FakeBlockPuller) PullBlockCalls(stub func(uint64) *common.Block) {
	fake.pullBlockMutex.Lock()
	defer fake.pullBlockMutex.Unlock()
	fake.PullBlockStub = stub
}

func (fake *FakeBlockPuller) PullBlockArgsForCall(i int) uint64 {
	fake.pullBlockMutex.RLock()
	defer fake.pullBlockMutex.RUnlock()
	argsForCall := fake.pullBlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockPuller) PullBlockReturns(result1 *common.Block) {
	fake.pullBlockMutex.Lock()
	defer fake.pullBlockMutex.Unlock()
	fake.PullBlockStub = nil
	fake.pullBlockReturns = struct {
		result1 *common.Block
	}{result1}
}

func (fake *FakeBlockPuller) PullBlockReturnsOnCall(i int, result1 *common.Block) {
	fake.pullBlockMutex.Lock()
	defer fake.pullBlockMutex.Unlock()
	fake.PullBlockStub = nil
	if fake.pullBlockReturnsOnCall == nil {
		fake.pullBlockReturnsOnCall = make(map[int]struct {
			result1 *common.Block
		})
	}
	fake.pullBlockReturnsOnCall[i] = struct {
		result1 *common.Block
	}{result1}
}

func (fake *FakeBlockPuller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.heightsByEndpointsMutex.RLock()
	defer fake.heightsByEndpointsMutex.RUnlock()
	fake.pullBlockMutex.RLock()
	defer fake.pullBlockMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockPuller) recordInvocation(key string, args []interface{}) {
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

var _ smartbft.BlockPuller = new(FakeBlockPuller)
