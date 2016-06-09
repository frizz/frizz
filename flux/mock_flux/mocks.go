// Automatically generated by MockGen. DO NOT EDIT!
// Source: kego.io/flux (interfaces: DispatcherInterface)

package mock_flux

import (
	gomock "github.com/golang/mock/gomock"
	flux "kego.io/flux"
)

// Mock of DispatcherInterface interface
type MockDispatcherInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockDispatcherInterfaceRecorder
}

// Recorder for MockDispatcherInterface (not exported)
type _MockDispatcherInterfaceRecorder struct {
	mock *MockDispatcherInterface
}

func NewMockDispatcherInterface(ctrl *gomock.Controller) *MockDispatcherInterface {
	mock := &MockDispatcherInterface{ctrl: ctrl}
	mock.recorder = &_MockDispatcherInterfaceRecorder{mock}
	return mock
}

func (_m *MockDispatcherInterface) EXPECT() *_MockDispatcherInterfaceRecorder {
	return _m.recorder
}

func (_m *MockDispatcherInterface) Dispatch(_param0 flux.ActionInterface) chan struct{} {
	ret := _m.ctrl.Call(_m, "Dispatch", _param0)
	ret0, _ := ret[0].(chan struct{})
	return ret0
}

func (_mr *_MockDispatcherInterfaceRecorder) Dispatch(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Dispatch", arg0)
}