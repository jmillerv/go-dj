// Code generated by MockGen. DO NOT EDIT.
// Source: ./media.go

// Package content is a generated GoMock package.
package content

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMedia is a mock of Media interface.
type MockMedia struct {
	ctrl     *gomock.Controller
	recorder *MockMediaMockRecorder
}

// MockMediaMockRecorder is the mock recorder for MockMedia.
type MockMediaMockRecorder struct {
	mock *MockMedia
}

// NewMockMedia creates a new mock instance.
func NewMockMedia(ctrl *gomock.Controller) *MockMedia {
	mock := &MockMedia{ctrl: ctrl}
	mock.recorder = &MockMediaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMedia) EXPECT() *MockMediaMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockMedia) Get() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockMediaMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMedia)(nil).Get))
}

// Play mocks base method.
func (m *MockMedia) Play() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Play")
	ret0, _ := ret[0].(error)
	return ret0
}

// Play indicates an expected call of Play.
func (mr *MockMediaMockRecorder) Play() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Play", reflect.TypeOf((*MockMedia)(nil).Play))
}

// Stop mocks base method.
func (m *MockMedia) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockMediaMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockMedia)(nil).Stop))
}