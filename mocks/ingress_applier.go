// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/seibert-media/k8s-ingress/model"
)

type DomainApplier struct {
	ApplyStub        func([]model.Domain) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 []model.Domain
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DomainApplier) Apply(arg1 []model.Domain) error {
	var arg1Copy []model.Domain
	if arg1 != nil {
		arg1Copy = make([]model.Domain, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 []model.Domain
	}{arg1Copy})
	fake.recordInvocation("Apply", []interface{}{arg1Copy})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.applyReturns.result1
}

func (fake *DomainApplier) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *DomainApplier) ApplyArgsForCall(i int) []model.Domain {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.applyArgsForCall[i].arg1
}

func (fake *DomainApplier) ApplyReturns(result1 error) {
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *DomainApplier) ApplyReturnsOnCall(i int, result1 error) {
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DomainApplier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DomainApplier) recordInvocation(key string, args []interface{}) {
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
