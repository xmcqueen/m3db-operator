// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3db-operator/pkg/m3admin/namespace/types.go

// Package namespace is a generated GoMock package.
package namespace

import (
	"reflect"

	"github.com/m3db/m3/src/query/generated/proto/admin"

	"github.com/golang/mock/gomock"
)

// MockNamespace is a mock of Namespace interface
type MockNamespace struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceMockRecorder
}

// MockNamespaceMockRecorder is the mock recorder for MockNamespace
type MockNamespaceMockRecorder struct {
	mock *MockNamespace
}

// NewMockNamespace creates a new mock instance
func NewMockNamespace(ctrl *gomock.Controller) *MockNamespace {
	mock := &MockNamespace{ctrl: ctrl}
	mock.recorder = &MockNamespaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespace) EXPECT() *MockNamespaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockNamespace) Create(namespace string) error {
	ret := m.ctrl.Call(m, "Create", namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockNamespaceMockRecorder) Create(namespace interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNamespace)(nil).Create), namespace)
}

// Get mocks base method
func (m *MockNamespace) Get() (*admin.NamespaceGetResponse, error) {
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*admin.NamespaceGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockNamespaceMockRecorder) Get() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNamespace)(nil).Get))
}

// Delete mocks base method
func (m *MockNamespace) Delete(namespace string) error {
	ret := m.ctrl.Call(m, "Delete", namespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockNamespaceMockRecorder) Delete(namespace interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNamespace)(nil).Delete), namespace)
}
