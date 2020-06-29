// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/interface.go

// Package pkg is a generated GoMock package.
package pkg

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/nuts-foundation/nuts-network/pkg/model"
	io "io"
	reflect "reflect"
	time "time"
)

// MockNetworkClient is a mock of NetworkClient interface
type MockNetworkClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkClientMockRecorder
}

// MockNetworkClientMockRecorder is the mock recorder for MockNetworkClient
type MockNetworkClientMockRecorder struct {
	mock *MockNetworkClient
}

// NewMockNetworkClient creates a new mock instance
func NewMockNetworkClient(ctrl *gomock.Controller) *MockNetworkClient {
	mock := &MockNetworkClient{ctrl: ctrl}
	mock.recorder = &MockNetworkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkClient) EXPECT() *MockNetworkClientMockRecorder {
	return m.recorder
}

// GetDocumentContents mocks base method
func (m *MockNetworkClient) GetDocumentContents(hash model.Hash) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocumentContents", hash)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocumentContents indicates an expected call of GetDocumentContents
func (mr *MockNetworkClientMockRecorder) GetDocumentContents(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocumentContents", reflect.TypeOf((*MockNetworkClient)(nil).GetDocumentContents), hash)
}

// GetDocument mocks base method
func (m *MockNetworkClient) GetDocument(hash model.Hash) (*model.DocumentDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDocument", hash)
	ret0, _ := ret[0].(*model.DocumentDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDocument indicates an expected call of GetDocument
func (mr *MockNetworkClientMockRecorder) GetDocument(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDocument", reflect.TypeOf((*MockNetworkClient)(nil).GetDocument), hash)
}

// AddDocumentWithContents mocks base method
func (m *MockNetworkClient) AddDocumentWithContents(timestamp time.Time, docType string, contents []byte) (*model.Document, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDocumentWithContents", timestamp, docType, contents)
	ret0, _ := ret[0].(*model.Document)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDocumentWithContents indicates an expected call of AddDocumentWithContents
func (mr *MockNetworkClientMockRecorder) AddDocumentWithContents(timestamp, docType, contents interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDocumentWithContents", reflect.TypeOf((*MockNetworkClient)(nil).AddDocumentWithContents), timestamp, docType, contents)
}

// ListDocuments mocks base method
func (m *MockNetworkClient) ListDocuments() ([]model.DocumentDescriptor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDocuments")
	ret0, _ := ret[0].([]model.DocumentDescriptor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDocuments indicates an expected call of ListDocuments
func (mr *MockNetworkClientMockRecorder) ListDocuments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDocuments", reflect.TypeOf((*MockNetworkClient)(nil).ListDocuments))
}
