// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotalservices/cf-mgmt/uaa"
)

type FakeManager struct {
	ListUsersStub        func() (map[string]string, error)
	listUsersMutex       sync.RWMutex
	listUsersArgsForCall []struct{}
	listUsersReturns     struct {
		result1 map[string]string
		result2 error
	}
	UsersByIDStub        func() (map[string]uaa.User, error)
	usersByIDMutex       sync.RWMutex
	usersByIDArgsForCall []struct{}
	usersByIDReturns     struct {
		result1 map[string]uaa.User
		result2 error
	}
	CreateExternalUserStub        func(userName, userEmail, externalID, origin string) (err error)
	createExternalUserMutex       sync.RWMutex
	createExternalUserArgsForCall []struct {
		userName   string
		userEmail  string
		externalID string
		origin     string
	}
	createExternalUserReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) ListUsers() (map[string]string, error) {
	fake.listUsersMutex.Lock()
	fake.listUsersArgsForCall = append(fake.listUsersArgsForCall, struct{}{})
	fake.recordInvocation("ListUsers", []interface{}{})
	fake.listUsersMutex.Unlock()
	if fake.ListUsersStub != nil {
		return fake.ListUsersStub()
	} else {
		return fake.listUsersReturns.result1, fake.listUsersReturns.result2
	}
}

func (fake *FakeManager) ListUsersCallCount() int {
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	return len(fake.listUsersArgsForCall)
}

func (fake *FakeManager) ListUsersReturns(result1 map[string]string, result2 error) {
	fake.ListUsersStub = nil
	fake.listUsersReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) UsersByID() (map[string]uaa.User, error) {
	fake.usersByIDMutex.Lock()
	fake.usersByIDArgsForCall = append(fake.usersByIDArgsForCall, struct{}{})
	fake.recordInvocation("UsersByID", []interface{}{})
	fake.usersByIDMutex.Unlock()
	if fake.UsersByIDStub != nil {
		return fake.UsersByIDStub()
	} else {
		return fake.usersByIDReturns.result1, fake.usersByIDReturns.result2
	}
}

func (fake *FakeManager) UsersByIDCallCount() int {
	fake.usersByIDMutex.RLock()
	defer fake.usersByIDMutex.RUnlock()
	return len(fake.usersByIDArgsForCall)
}

func (fake *FakeManager) UsersByIDReturns(result1 map[string]uaa.User, result2 error) {
	fake.UsersByIDStub = nil
	fake.usersByIDReturns = struct {
		result1 map[string]uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) CreateExternalUser(userName string, userEmail string, externalID string, origin string) (err error) {
	fake.createExternalUserMutex.Lock()
	fake.createExternalUserArgsForCall = append(fake.createExternalUserArgsForCall, struct {
		userName   string
		userEmail  string
		externalID string
		origin     string
	}{userName, userEmail, externalID, origin})
	fake.recordInvocation("CreateExternalUser", []interface{}{userName, userEmail, externalID, origin})
	fake.createExternalUserMutex.Unlock()
	if fake.CreateExternalUserStub != nil {
		return fake.CreateExternalUserStub(userName, userEmail, externalID, origin)
	} else {
		return fake.createExternalUserReturns.result1
	}
}

func (fake *FakeManager) CreateExternalUserCallCount() int {
	fake.createExternalUserMutex.RLock()
	defer fake.createExternalUserMutex.RUnlock()
	return len(fake.createExternalUserArgsForCall)
}

func (fake *FakeManager) CreateExternalUserArgsForCall(i int) (string, string, string, string) {
	fake.createExternalUserMutex.RLock()
	defer fake.createExternalUserMutex.RUnlock()
	return fake.createExternalUserArgsForCall[i].userName, fake.createExternalUserArgsForCall[i].userEmail, fake.createExternalUserArgsForCall[i].externalID, fake.createExternalUserArgsForCall[i].origin
}

func (fake *FakeManager) CreateExternalUserReturns(result1 error) {
	fake.CreateExternalUserStub = nil
	fake.createExternalUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	fake.usersByIDMutex.RLock()
	defer fake.usersByIDMutex.RUnlock()
	fake.createExternalUserMutex.RLock()
	defer fake.createExternalUserMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ uaa.Manager = new(FakeManager)
