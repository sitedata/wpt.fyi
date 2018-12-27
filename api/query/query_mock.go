// Code generated by MockGen. DO NOT EDIT.
// Source: api/query/query.go

// Package query is a generated GoMock package.
package query

import (
	http "net/http"
	reflect "reflect"
	time "time"

	golang_set "github.com/deckarep/golang-set"
	gomock "github.com/golang/mock/gomock"
	shared "github.com/web-platform-tests/wpt.fyi/shared"
)

// MocksharedInterface is a mock of sharedInterface interface
type MocksharedInterface struct {
	ctrl     *gomock.Controller
	recorder *MocksharedInterfaceMockRecorder
}

// MocksharedInterfaceMockRecorder is the mock recorder for MocksharedInterface
type MocksharedInterfaceMockRecorder struct {
	mock *MocksharedInterface
}

// NewMocksharedInterface creates a new mock instance
func NewMocksharedInterface(ctrl *gomock.Controller) *MocksharedInterface {
	mock := &MocksharedInterface{ctrl: ctrl}
	mock.recorder = &MocksharedInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocksharedInterface) EXPECT() *MocksharedInterfaceMockRecorder {
	return m.recorder
}

// ParseQueryParamInt mocks base method
func (m *MocksharedInterface) ParseQueryParamInt(r *http.Request, key string) (*int, error) {
	ret := m.ctrl.Call(m, "ParseQueryParamInt", r, key)
	ret0, _ := ret[0].(*int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseQueryParamInt indicates an expected call of ParseQueryParamInt
func (mr *MocksharedInterfaceMockRecorder) ParseQueryParamInt(r, key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseQueryParamInt", reflect.TypeOf((*MocksharedInterface)(nil).ParseQueryParamInt), r, key)
}

// ParseQueryFilterParams mocks base method
func (m *MocksharedInterface) ParseQueryFilterParams(arg0 *http.Request) (shared.QueryFilter, error) {
	ret := m.ctrl.Call(m, "ParseQueryFilterParams", arg0)
	ret0, _ := ret[0].(shared.QueryFilter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseQueryFilterParams indicates an expected call of ParseQueryFilterParams
func (mr *MocksharedInterfaceMockRecorder) ParseQueryFilterParams(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseQueryFilterParams", reflect.TypeOf((*MocksharedInterface)(nil).ParseQueryFilterParams), arg0)
}

// LoadTestRuns mocks base method
func (m *MocksharedInterface) LoadTestRuns(arg0 shared.ProductSpecs, arg1 golang_set.Set, arg2 shared.SHAs, arg3, arg4 *time.Time, arg5, arg6 *int) (shared.TestRunsByProduct, error) {
	ret := m.ctrl.Call(m, "LoadTestRuns", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(shared.TestRunsByProduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTestRuns indicates an expected call of LoadTestRuns
func (mr *MocksharedInterfaceMockRecorder) LoadTestRuns(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTestRuns", reflect.TypeOf((*MocksharedInterface)(nil).LoadTestRuns), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// LoadTestRunsByIDs mocks base method
func (m *MocksharedInterface) LoadTestRunsByIDs(ids shared.TestRunIDs) (shared.TestRuns, error) {
	ret := m.ctrl.Call(m, "LoadTestRunsByIDs", ids)
	ret0, _ := ret[0].(shared.TestRuns)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTestRunsByIDs indicates an expected call of LoadTestRunsByIDs
func (mr *MocksharedInterfaceMockRecorder) LoadTestRunsByIDs(ids interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTestRunsByIDs", reflect.TypeOf((*MocksharedInterface)(nil).LoadTestRunsByIDs), ids)
}

// LoadTestRun mocks base method
func (m *MocksharedInterface) LoadTestRun(arg0 int64) (*shared.TestRun, error) {
	ret := m.ctrl.Call(m, "LoadTestRun", arg0)
	ret0, _ := ret[0].(*shared.TestRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTestRun indicates an expected call of LoadTestRun
func (mr *MocksharedInterfaceMockRecorder) LoadTestRun(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTestRun", reflect.TypeOf((*MocksharedInterface)(nil).LoadTestRun), arg0)
}
