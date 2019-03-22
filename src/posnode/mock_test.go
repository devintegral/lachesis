// Code generated by MockGen. DO NOT EDIT.
// Source: consensus.go

// Package posnode is a generated GoMock package.
package posnode

import (
	common "github.com/Fantom-foundation/go-lachesis/src/common"
	wire "github.com/Fantom-foundation/go-lachesis/src/posnode/wire"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConsensus is a mock of Consensus interface
type MockConsensus struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusMockRecorder
}

// MockConsensusMockRecorder is the mock recorder for MockConsensus
type MockConsensusMockRecorder struct {
	mock *MockConsensus
}

// NewMockConsensus creates a new mock instance
func NewMockConsensus(ctrl *gomock.Controller) *MockConsensus {
	mock := &MockConsensus{ctrl: ctrl}
	mock.recorder = &MockConsensusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsensus) EXPECT() *MockConsensusMockRecorder {
	return m.recorder
}

// PushEvent mocks base method
func (m *MockConsensus) PushEvent(e *wire.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PushEvent", e)
}

// PushEvent indicates an expected call of PushEvent
func (mr *MockConsensusMockRecorder) PushEvent(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushEvent", reflect.TypeOf((*MockConsensus)(nil).PushEvent), e)
}

// GetEvent mocks base method
func (m *MockConsensus) GetEvent(creator common.Address, index uint64) *wire.Event {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", creator, index)
	ret0, _ := ret[0].(*wire.Event)
	return ret0
}

// GetEvent indicates an expected call of GetEvent
func (mr *MockConsensusMockRecorder) GetEvent(creator, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockConsensus)(nil).GetEvent), creator, index)
}

// LastKnownEvent mocks base method
func (m *MockConsensus) LastKnownEvent(creator common.Address) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastKnownEvent", creator)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// LastKnownEvent indicates an expected call of LastKnownEvent
func (mr *MockConsensusMockRecorder) LastKnownEvent(creator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastKnownEvent", reflect.TypeOf((*MockConsensus)(nil).LastKnownEvent), creator)
}