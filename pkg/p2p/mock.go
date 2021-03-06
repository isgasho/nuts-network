// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/p2p/interface.go

// Package p2p is a generated GoMock package.
package p2p

import (
	gomock "github.com/golang/mock/gomock"
	network "github.com/nuts-foundation/nuts-network/network"
	model "github.com/nuts-foundation/nuts-network/pkg/model"
	stats "github.com/nuts-foundation/nuts-network/pkg/stats"
	reflect "reflect"
)

// MockP2PNetwork is a mock of P2PNetwork interface
type MockP2PNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockP2PNetworkMockRecorder
}

// MockP2PNetworkMockRecorder is the mock recorder for MockP2PNetwork
type MockP2PNetworkMockRecorder struct {
	mock *MockP2PNetwork
}

// NewMockP2PNetwork creates a new mock instance
func NewMockP2PNetwork(ctrl *gomock.Controller) *MockP2PNetwork {
	mock := &MockP2PNetwork{ctrl: ctrl}
	mock.recorder = &MockP2PNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockP2PNetwork) EXPECT() *MockP2PNetworkMockRecorder {
	return m.recorder
}

// Statistics mocks base method
func (m *MockP2PNetwork) Statistics() []stats.Statistic {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Statistics")
	ret0, _ := ret[0].([]stats.Statistic)
	return ret0
}

// Statistics indicates an expected call of Statistics
func (mr *MockP2PNetworkMockRecorder) Statistics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statistics", reflect.TypeOf((*MockP2PNetwork)(nil).Statistics))
}

// Configure mocks base method
func (m *MockP2PNetwork) Configure(config P2PNetworkConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configure", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// Configure indicates an expected call of Configure
func (mr *MockP2PNetworkMockRecorder) Configure(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockP2PNetwork)(nil).Configure), config)
}

// Configured mocks base method
func (m *MockP2PNetwork) Configured() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Configured")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Configured indicates an expected call of Configured
func (mr *MockP2PNetworkMockRecorder) Configured() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configured", reflect.TypeOf((*MockP2PNetwork)(nil).Configured))
}

// Start mocks base method
func (m *MockP2PNetwork) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockP2PNetworkMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockP2PNetwork)(nil).Start))
}

// Stop mocks base method
func (m *MockP2PNetwork) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockP2PNetworkMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockP2PNetwork)(nil).Stop))
}

// AddRemoteNode mocks base method
func (m *MockP2PNetwork) AddRemoteNode(node model.NodeInfo) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRemoteNode", node)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AddRemoteNode indicates an expected call of AddRemoteNode
func (mr *MockP2PNetworkMockRecorder) AddRemoteNode(node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteNode", reflect.TypeOf((*MockP2PNetwork)(nil).AddRemoteNode), node)
}

// ReceivedMessages mocks base method
func (m *MockP2PNetwork) ReceivedMessages() MessageQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivedMessages")
	ret0, _ := ret[0].(MessageQueue)
	return ret0
}

// ReceivedMessages indicates an expected call of ReceivedMessages
func (mr *MockP2PNetworkMockRecorder) ReceivedMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedMessages", reflect.TypeOf((*MockP2PNetwork)(nil).ReceivedMessages))
}

// Send mocks base method
func (m *MockP2PNetwork) Send(peer model.PeerID, message *network.NetworkMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", peer, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockP2PNetworkMockRecorder) Send(peer, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockP2PNetwork)(nil).Send), peer, message)
}

// Broadcast mocks base method
func (m *MockP2PNetwork) Broadcast(message *network.NetworkMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Broadcast", message)
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockP2PNetworkMockRecorder) Broadcast(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockP2PNetwork)(nil).Broadcast), message)
}

// Peers mocks base method
func (m *MockP2PNetwork) Peers() []Peer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers")
	ret0, _ := ret[0].([]Peer)
	return ret0
}

// Peers indicates an expected call of Peers
func (mr *MockP2PNetworkMockRecorder) Peers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockP2PNetwork)(nil).Peers))
}

// MockMessageQueue is a mock of MessageQueue interface
type MockMessageQueue struct {
	ctrl     *gomock.Controller
	recorder *MockMessageQueueMockRecorder
}

// MockMessageQueueMockRecorder is the mock recorder for MockMessageQueue
type MockMessageQueueMockRecorder struct {
	mock *MockMessageQueue
}

// NewMockMessageQueue creates a new mock instance
func NewMockMessageQueue(ctrl *gomock.Controller) *MockMessageQueue {
	mock := &MockMessageQueue{ctrl: ctrl}
	mock.recorder = &MockMessageQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageQueue) EXPECT() *MockMessageQueueMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockMessageQueue) Get() PeerMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(PeerMessage)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockMessageQueueMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMessageQueue)(nil).Get))
}
