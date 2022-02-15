// Code generated by MockGen. DO NOT EDIT.
// Source: cost_estimate.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockCostEstimates is a mock of CostEstimates interface.
type MockCostEstimates struct {
	ctrl     *gomock.Controller
	recorder *MockCostEstimatesMockRecorder
}

// MockCostEstimatesMockRecorder is the mock recorder for MockCostEstimates.
type MockCostEstimatesMockRecorder struct {
	mock *MockCostEstimates
}

// NewMockCostEstimates creates a new mock instance.
func NewMockCostEstimates(ctrl *gomock.Controller) *MockCostEstimates {
	mock := &MockCostEstimates{ctrl: ctrl}
	mock.recorder = &MockCostEstimatesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCostEstimates) EXPECT() *MockCostEstimatesMockRecorder {
	return m.recorder
}

// Logs mocks base method.
func (m *MockCostEstimates) Logs(ctx context.Context, costEstimateID string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", ctx, costEstimateID)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Logs indicates an expected call of Logs.
func (mr *MockCostEstimatesMockRecorder) Logs(ctx, costEstimateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*MockCostEstimates)(nil).Logs), ctx, costEstimateID)
}

// Read mocks base method.
func (m *MockCostEstimates) Read(ctx context.Context, costEstimateID string) (*tfe.CostEstimate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, costEstimateID)
	ret0, _ := ret[0].(*tfe.CostEstimate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockCostEstimatesMockRecorder) Read(ctx, costEstimateID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCostEstimates)(nil).Read), ctx, costEstimateID)
}
