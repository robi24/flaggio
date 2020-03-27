// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/victorkt/flaggio/internal/flaggio (interfaces: Evaluator,Identifier)

// Package flaggio_mock is a generated GoMock package.
package flaggio_mock

import (
	gomock "github.com/golang/mock/gomock"
	flaggio "github.com/victorkt/flaggio/internal/flaggio"
	reflect "reflect"
)

// MockEvaluator is a mock of Evaluator interface
type MockEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockEvaluatorMockRecorder
}

// MockEvaluatorMockRecorder is the mock recorder for MockEvaluator
type MockEvaluatorMockRecorder struct {
	mock *MockEvaluator
}

// NewMockEvaluator creates a new mock instance
func NewMockEvaluator(ctrl *gomock.Controller) *MockEvaluator {
	mock := &MockEvaluator{ctrl: ctrl}
	mock.recorder = &MockEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEvaluator) EXPECT() *MockEvaluatorMockRecorder {
	return m.recorder
}

// Evaluate mocks base method
func (m *MockEvaluator) Evaluate(arg0 map[string]interface{}) (flaggio.EvalResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Evaluate", arg0)
	ret0, _ := ret[0].(flaggio.EvalResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Evaluate indicates an expected call of Evaluate
func (mr *MockEvaluatorMockRecorder) Evaluate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Evaluate", reflect.TypeOf((*MockEvaluator)(nil).Evaluate), arg0)
}

// MockIdentifier is a mock of Identifier interface
type MockIdentifier struct {
	ctrl     *gomock.Controller
	recorder *MockIdentifierMockRecorder
}

// MockIdentifierMockRecorder is the mock recorder for MockIdentifier
type MockIdentifierMockRecorder struct {
	mock *MockIdentifier
}

// NewMockIdentifier creates a new mock instance
func NewMockIdentifier(ctrl *gomock.Controller) *MockIdentifier {
	mock := &MockIdentifier{ctrl: ctrl}
	mock.recorder = &MockIdentifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIdentifier) EXPECT() *MockIdentifierMockRecorder {
	return m.recorder
}

// GetID mocks base method
func (m *MockIdentifier) GetID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockIdentifierMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockIdentifier)(nil).GetID))
}
