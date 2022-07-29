// Code generated by MockGen. DO NOT EDIT.
// Source: ../manager.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "d7y.io/api/pkg/apis/manager/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockManagerClient is a mock of ManagerClient interface.
type MockManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagerClientMockRecorder
}

// MockManagerClientMockRecorder is the mock recorder for MockManagerClient.
type MockManagerClientMockRecorder struct {
	mock *MockManagerClient
}

// NewMockManagerClient creates a new mock instance.
func NewMockManagerClient(ctrl *gomock.Controller) *MockManagerClient {
	mock := &MockManagerClient{ctrl: ctrl}
	mock.recorder = &MockManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerClient) EXPECT() *MockManagerClientMockRecorder {
	return m.recorder
}

// GetObjectStorage mocks base method.
func (m *MockManagerClient) GetObjectStorage(ctx context.Context, in *v1.GetObjectStorageRequest, opts ...grpc.CallOption) (*v1.ObjectStorage, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectStorage", varargs...)
	ret0, _ := ret[0].(*v1.ObjectStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectStorage indicates an expected call of GetObjectStorage.
func (mr *MockManagerClientMockRecorder) GetObjectStorage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStorage", reflect.TypeOf((*MockManagerClient)(nil).GetObjectStorage), varargs...)
}

// GetScheduler mocks base method.
func (m *MockManagerClient) GetScheduler(ctx context.Context, in *v1.GetSchedulerRequest, opts ...grpc.CallOption) (*v1.Scheduler, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetScheduler", varargs...)
	ret0, _ := ret[0].(*v1.Scheduler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduler indicates an expected call of GetScheduler.
func (mr *MockManagerClientMockRecorder) GetScheduler(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduler", reflect.TypeOf((*MockManagerClient)(nil).GetScheduler), varargs...)
}

// GetSeedPeer mocks base method.
func (m *MockManagerClient) GetSeedPeer(ctx context.Context, in *v1.GetSeedPeerRequest, opts ...grpc.CallOption) (*v1.SeedPeer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSeedPeer", varargs...)
	ret0, _ := ret[0].(*v1.SeedPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeedPeer indicates an expected call of GetSeedPeer.
func (mr *MockManagerClientMockRecorder) GetSeedPeer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeedPeer", reflect.TypeOf((*MockManagerClient)(nil).GetSeedPeer), varargs...)
}

// KeepAlive mocks base method.
func (m *MockManagerClient) KeepAlive(ctx context.Context, opts ...grpc.CallOption) (v1.Manager_KeepAliveClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "KeepAlive", varargs...)
	ret0, _ := ret[0].(v1.Manager_KeepAliveClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeepAlive indicates an expected call of KeepAlive.
func (mr *MockManagerClientMockRecorder) KeepAlive(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepAlive", reflect.TypeOf((*MockManagerClient)(nil).KeepAlive), varargs...)
}

// ListBuckets mocks base method.
func (m *MockManagerClient) ListBuckets(ctx context.Context, in *v1.ListBucketsRequest, opts ...grpc.CallOption) (*v1.ListBucketsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuckets", varargs...)
	ret0, _ := ret[0].(*v1.ListBucketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockManagerClientMockRecorder) ListBuckets(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockManagerClient)(nil).ListBuckets), varargs...)
}

// ListSchedulers mocks base method.
func (m *MockManagerClient) ListSchedulers(ctx context.Context, in *v1.ListSchedulersRequest, opts ...grpc.CallOption) (*v1.ListSchedulersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchedulers", varargs...)
	ret0, _ := ret[0].(*v1.ListSchedulersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulers indicates an expected call of ListSchedulers.
func (mr *MockManagerClientMockRecorder) ListSchedulers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulers", reflect.TypeOf((*MockManagerClient)(nil).ListSchedulers), varargs...)
}

// UpdateScheduler mocks base method.
func (m *MockManagerClient) UpdateScheduler(ctx context.Context, in *v1.UpdateSchedulerRequest, opts ...grpc.CallOption) (*v1.Scheduler, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateScheduler", varargs...)
	ret0, _ := ret[0].(*v1.Scheduler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduler indicates an expected call of UpdateScheduler.
func (mr *MockManagerClientMockRecorder) UpdateScheduler(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduler", reflect.TypeOf((*MockManagerClient)(nil).UpdateScheduler), varargs...)
}

// UpdateSeedPeer mocks base method.
func (m *MockManagerClient) UpdateSeedPeer(ctx context.Context, in *v1.UpdateSeedPeerRequest, opts ...grpc.CallOption) (*v1.SeedPeer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSeedPeer", varargs...)
	ret0, _ := ret[0].(*v1.SeedPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSeedPeer indicates an expected call of UpdateSeedPeer.
func (mr *MockManagerClientMockRecorder) UpdateSeedPeer(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSeedPeer", reflect.TypeOf((*MockManagerClient)(nil).UpdateSeedPeer), varargs...)
}

// MockManager_KeepAliveClient is a mock of Manager_KeepAliveClient interface.
type MockManager_KeepAliveClient struct {
	ctrl     *gomock.Controller
	recorder *MockManager_KeepAliveClientMockRecorder
}

// MockManager_KeepAliveClientMockRecorder is the mock recorder for MockManager_KeepAliveClient.
type MockManager_KeepAliveClientMockRecorder struct {
	mock *MockManager_KeepAliveClient
}

// NewMockManager_KeepAliveClient creates a new mock instance.
func NewMockManager_KeepAliveClient(ctrl *gomock.Controller) *MockManager_KeepAliveClient {
	mock := &MockManager_KeepAliveClient{ctrl: ctrl}
	mock.recorder = &MockManager_KeepAliveClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager_KeepAliveClient) EXPECT() *MockManager_KeepAliveClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockManager_KeepAliveClient) CloseAndRecv() (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockManager_KeepAliveClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockManager_KeepAliveClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockManager_KeepAliveClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockManager_KeepAliveClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockManager_KeepAliveClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).Context))
}

// Header mocks base method.
func (m *MockManager_KeepAliveClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockManager_KeepAliveClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m_2 *MockManager_KeepAliveClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockManager_KeepAliveClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockManager_KeepAliveClient) Send(arg0 *v1.KeepAliveRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockManager_KeepAliveClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockManager_KeepAliveClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockManager_KeepAliveClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockManager_KeepAliveClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockManager_KeepAliveClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockManager_KeepAliveClient)(nil).Trailer))
}

// MockManagerServer is a mock of ManagerServer interface.
type MockManagerServer struct {
	ctrl     *gomock.Controller
	recorder *MockManagerServerMockRecorder
}

// MockManagerServerMockRecorder is the mock recorder for MockManagerServer.
type MockManagerServerMockRecorder struct {
	mock *MockManagerServer
}

// NewMockManagerServer creates a new mock instance.
func NewMockManagerServer(ctrl *gomock.Controller) *MockManagerServer {
	mock := &MockManagerServer{ctrl: ctrl}
	mock.recorder = &MockManagerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerServer) EXPECT() *MockManagerServerMockRecorder {
	return m.recorder
}

// GetObjectStorage mocks base method.
func (m *MockManagerServer) GetObjectStorage(arg0 context.Context, arg1 *v1.GetObjectStorageRequest) (*v1.ObjectStorage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectStorage", arg0, arg1)
	ret0, _ := ret[0].(*v1.ObjectStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectStorage indicates an expected call of GetObjectStorage.
func (mr *MockManagerServerMockRecorder) GetObjectStorage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectStorage", reflect.TypeOf((*MockManagerServer)(nil).GetObjectStorage), arg0, arg1)
}

// GetScheduler mocks base method.
func (m *MockManagerServer) GetScheduler(arg0 context.Context, arg1 *v1.GetSchedulerRequest) (*v1.Scheduler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScheduler", arg0, arg1)
	ret0, _ := ret[0].(*v1.Scheduler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScheduler indicates an expected call of GetScheduler.
func (mr *MockManagerServerMockRecorder) GetScheduler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScheduler", reflect.TypeOf((*MockManagerServer)(nil).GetScheduler), arg0, arg1)
}

// GetSeedPeer mocks base method.
func (m *MockManagerServer) GetSeedPeer(arg0 context.Context, arg1 *v1.GetSeedPeerRequest) (*v1.SeedPeer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeedPeer", arg0, arg1)
	ret0, _ := ret[0].(*v1.SeedPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSeedPeer indicates an expected call of GetSeedPeer.
func (mr *MockManagerServerMockRecorder) GetSeedPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeedPeer", reflect.TypeOf((*MockManagerServer)(nil).GetSeedPeer), arg0, arg1)
}

// KeepAlive mocks base method.
func (m *MockManagerServer) KeepAlive(arg0 v1.Manager_KeepAliveServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeepAlive", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// KeepAlive indicates an expected call of KeepAlive.
func (mr *MockManagerServerMockRecorder) KeepAlive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepAlive", reflect.TypeOf((*MockManagerServer)(nil).KeepAlive), arg0)
}

// ListBuckets mocks base method.
func (m *MockManagerServer) ListBuckets(arg0 context.Context, arg1 *v1.ListBucketsRequest) (*v1.ListBucketsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuckets", arg0, arg1)
	ret0, _ := ret[0].(*v1.ListBucketsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockManagerServerMockRecorder) ListBuckets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockManagerServer)(nil).ListBuckets), arg0, arg1)
}

// ListSchedulers mocks base method.
func (m *MockManagerServer) ListSchedulers(arg0 context.Context, arg1 *v1.ListSchedulersRequest) (*v1.ListSchedulersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedulers", arg0, arg1)
	ret0, _ := ret[0].(*v1.ListSchedulersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedulers indicates an expected call of ListSchedulers.
func (mr *MockManagerServerMockRecorder) ListSchedulers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedulers", reflect.TypeOf((*MockManagerServer)(nil).ListSchedulers), arg0, arg1)
}

// UpdateScheduler mocks base method.
func (m *MockManagerServer) UpdateScheduler(arg0 context.Context, arg1 *v1.UpdateSchedulerRequest) (*v1.Scheduler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScheduler", arg0, arg1)
	ret0, _ := ret[0].(*v1.Scheduler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateScheduler indicates an expected call of UpdateScheduler.
func (mr *MockManagerServerMockRecorder) UpdateScheduler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScheduler", reflect.TypeOf((*MockManagerServer)(nil).UpdateScheduler), arg0, arg1)
}

// UpdateSeedPeer mocks base method.
func (m *MockManagerServer) UpdateSeedPeer(arg0 context.Context, arg1 *v1.UpdateSeedPeerRequest) (*v1.SeedPeer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSeedPeer", arg0, arg1)
	ret0, _ := ret[0].(*v1.SeedPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSeedPeer indicates an expected call of UpdateSeedPeer.
func (mr *MockManagerServerMockRecorder) UpdateSeedPeer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSeedPeer", reflect.TypeOf((*MockManagerServer)(nil).UpdateSeedPeer), arg0, arg1)
}

// MockManager_KeepAliveServer is a mock of Manager_KeepAliveServer interface.
type MockManager_KeepAliveServer struct {
	ctrl     *gomock.Controller
	recorder *MockManager_KeepAliveServerMockRecorder
}

// MockManager_KeepAliveServerMockRecorder is the mock recorder for MockManager_KeepAliveServer.
type MockManager_KeepAliveServerMockRecorder struct {
	mock *MockManager_KeepAliveServer
}

// NewMockManager_KeepAliveServer creates a new mock instance.
func NewMockManager_KeepAliveServer(ctrl *gomock.Controller) *MockManager_KeepAliveServer {
	mock := &MockManager_KeepAliveServer{ctrl: ctrl}
	mock.recorder = &MockManager_KeepAliveServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager_KeepAliveServer) EXPECT() *MockManager_KeepAliveServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockManager_KeepAliveServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockManager_KeepAliveServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockManager_KeepAliveServer) Recv() (*v1.KeepAliveRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.KeepAliveRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockManager_KeepAliveServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockManager_KeepAliveServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockManager_KeepAliveServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).RecvMsg), m)
}

// SendAndClose mocks base method.
func (m *MockManager_KeepAliveServer) SendAndClose(arg0 *emptypb.Empty) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockManager_KeepAliveServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).SendAndClose), arg0)
}

// SendHeader mocks base method.
func (m *MockManager_KeepAliveServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockManager_KeepAliveServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockManager_KeepAliveServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockManager_KeepAliveServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockManager_KeepAliveServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockManager_KeepAliveServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockManager_KeepAliveServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockManager_KeepAliveServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockManager_KeepAliveServer)(nil).SetTrailer), arg0)
}
