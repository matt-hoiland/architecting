// Code generated by MockGen. DO NOT EDIT.
// Source: dependencies.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockCollector is a mock of Collector interface.
type MockCollector struct {
	ctrl     *gomock.Controller
	recorder *MockCollectorMockRecorder
}

// MockCollectorMockRecorder is the mock recorder for MockCollector.
type MockCollectorMockRecorder struct {
	mock *MockCollector
}

// NewMockCollector creates a new mock instance.
func NewMockCollector(ctrl *gomock.Controller) *MockCollector {
	mock := &MockCollector{ctrl: ctrl}
	mock.recorder = &MockCollectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCollector) EXPECT() *MockCollectorMockRecorder {
	return m.recorder
}

// DeleteOne mocks base method.
func (m *MockCollector) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteOne", varargs...)
	ret0, _ := ret[0].(*mongo.DeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockCollectorMockRecorder) DeleteOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockCollector)(nil).DeleteOne), varargs...)
}

// FindOne mocks base method.
func (m *MockCollector) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCollectorMockRecorder) FindOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockCollector)(nil).FindOne), varargs...)
}

// InsertOne mocks base method.
func (m *MockCollector) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, document}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockCollectorMockRecorder) InsertOne(ctx, document interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, document}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockCollector)(nil).InsertOne), varargs...)
}

// UpdateByID mocks base method.
func (m *MockCollector) UpdateByID(ctx context.Context, id, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id, update}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateByID", varargs...)
	ret0, _ := ret[0].(*mongo.UpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockCollectorMockRecorder) UpdateByID(ctx, id, update interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id, update}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockCollector)(nil).UpdateByID), varargs...)
}