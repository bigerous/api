// Code generated by MockGen. DO NOT EDIT.
// Source: ../scheduler.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "d7y.io/api/pkg/apis/scheduler/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockisRegisterResult_DirectPiece is a mock of isRegisterResult_DirectPiece interface.
type MockisRegisterResult_DirectPiece struct {
	ctrl     *gomock.Controller
	recorder *MockisRegisterResult_DirectPieceMockRecorder
}

// MockisRegisterResult_DirectPieceMockRecorder is the mock recorder for MockisRegisterResult_DirectPiece.
type MockisRegisterResult_DirectPieceMockRecorder struct {
	mock *MockisRegisterResult_DirectPiece
}

// NewMockisRegisterResult_DirectPiece creates a new mock instance.
func NewMockisRegisterResult_DirectPiece(ctrl *gomock.Controller) *MockisRegisterResult_DirectPiece {
	mock := &MockisRegisterResult_DirectPiece{ctrl: ctrl}
	mock.recorder = &MockisRegisterResult_DirectPieceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisRegisterResult_DirectPiece) EXPECT() *MockisRegisterResult_DirectPieceMockRecorder {
	return m.recorder
}

// isRegisterResult_DirectPiece mocks base method.
func (m *MockisRegisterResult_DirectPiece) isRegisterResult_DirectPiece() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isRegisterResult_DirectPiece")
}

// isRegisterResult_DirectPiece indicates an expected call of isRegisterResult_DirectPiece.
func (mr *MockisRegisterResult_DirectPieceMockRecorder) isRegisterResult_DirectPiece() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isRegisterResult_DirectPiece", reflect.TypeOf((*MockisRegisterResult_DirectPiece)(nil).isRegisterResult_DirectPiece))
}

// MockisPeerPacket_Errordetails is a mock of isPeerPacket_Errordetails interface.
type MockisPeerPacket_Errordetails struct {
	ctrl     *gomock.Controller
	recorder *MockisPeerPacket_ErrordetailsMockRecorder
}

// MockisPeerPacket_ErrordetailsMockRecorder is the mock recorder for MockisPeerPacket_Errordetails.
type MockisPeerPacket_ErrordetailsMockRecorder struct {
	mock *MockisPeerPacket_Errordetails
}

// NewMockisPeerPacket_Errordetails creates a new mock instance.
func NewMockisPeerPacket_Errordetails(ctrl *gomock.Controller) *MockisPeerPacket_Errordetails {
	mock := &MockisPeerPacket_Errordetails{ctrl: ctrl}
	mock.recorder = &MockisPeerPacket_ErrordetailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisPeerPacket_Errordetails) EXPECT() *MockisPeerPacket_ErrordetailsMockRecorder {
	return m.recorder
}

// isPeerPacket_Errordetails mocks base method.
func (m *MockisPeerPacket_Errordetails) isPeerPacket_Errordetails() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isPeerPacket_Errordetails")
}

// isPeerPacket_Errordetails indicates an expected call of isPeerPacket_Errordetails.
func (mr *MockisPeerPacket_ErrordetailsMockRecorder) isPeerPacket_Errordetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isPeerPacket_Errordetails", reflect.TypeOf((*MockisPeerPacket_Errordetails)(nil).isPeerPacket_Errordetails))
}

// MockisPeerResult_Errordetails is a mock of isPeerResult_Errordetails interface.
type MockisPeerResult_Errordetails struct {
	ctrl     *gomock.Controller
	recorder *MockisPeerResult_ErrordetailsMockRecorder
}

// MockisPeerResult_ErrordetailsMockRecorder is the mock recorder for MockisPeerResult_Errordetails.
type MockisPeerResult_ErrordetailsMockRecorder struct {
	mock *MockisPeerResult_Errordetails
}

// NewMockisPeerResult_Errordetails creates a new mock instance.
func NewMockisPeerResult_Errordetails(ctrl *gomock.Controller) *MockisPeerResult_Errordetails {
	mock := &MockisPeerResult_Errordetails{ctrl: ctrl}
	mock.recorder = &MockisPeerResult_ErrordetailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisPeerResult_Errordetails) EXPECT() *MockisPeerResult_ErrordetailsMockRecorder {
	return m.recorder
}

// isPeerResult_Errordetails mocks base method.
func (m *MockisPeerResult_Errordetails) isPeerResult_Errordetails() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isPeerResult_Errordetails")
}

// isPeerResult_Errordetails indicates an expected call of isPeerResult_Errordetails.
func (mr *MockisPeerResult_ErrordetailsMockRecorder) isPeerResult_Errordetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isPeerResult_Errordetails", reflect.TypeOf((*MockisPeerResult_Errordetails)(nil).isPeerResult_Errordetails))
}

// MockSchedulerClient is a mock of SchedulerClient interface.
type MockSchedulerClient struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerClientMockRecorder
}

// MockSchedulerClientMockRecorder is the mock recorder for MockSchedulerClient.
type MockSchedulerClientMockRecorder struct {
	mock *MockSchedulerClient
}

// NewMockSchedulerClient creates a new mock instance.
func NewMockSchedulerClient(ctrl *gomock.Controller) *MockSchedulerClient {
	mock := &MockSchedulerClient{ctrl: ctrl}
	mock.recorder = &MockSchedulerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerClient) EXPECT() *MockSchedulerClientMockRecorder {
	return m.recorder
}

// AnnounceTask mocks base method.
func (m *MockSchedulerClient) AnnounceTask(ctx context.Context, in *v1.AnnounceTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AnnounceTask", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceTask indicates an expected call of AnnounceTask.
func (mr *MockSchedulerClientMockRecorder) AnnounceTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceTask", reflect.TypeOf((*MockSchedulerClient)(nil).AnnounceTask), varargs...)
}

// LeaveTask mocks base method.
func (m *MockSchedulerClient) LeaveTask(ctx context.Context, in *v1.PeerTarget, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LeaveTask", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaveTask indicates an expected call of LeaveTask.
func (mr *MockSchedulerClientMockRecorder) LeaveTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveTask", reflect.TypeOf((*MockSchedulerClient)(nil).LeaveTask), varargs...)
}

// RegisterPeerTask mocks base method.
func (m *MockSchedulerClient) RegisterPeerTask(ctx context.Context, in *v1.PeerTaskRequest, opts ...grpc.CallOption) (*v1.RegisterResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterPeerTask", varargs...)
	ret0, _ := ret[0].(*v1.RegisterResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterPeerTask indicates an expected call of RegisterPeerTask.
func (mr *MockSchedulerClientMockRecorder) RegisterPeerTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPeerTask", reflect.TypeOf((*MockSchedulerClient)(nil).RegisterPeerTask), varargs...)
}

// ReportPeerResult mocks base method.
func (m *MockSchedulerClient) ReportPeerResult(ctx context.Context, in *v1.PeerResult, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReportPeerResult", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportPeerResult indicates an expected call of ReportPeerResult.
func (mr *MockSchedulerClientMockRecorder) ReportPeerResult(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportPeerResult", reflect.TypeOf((*MockSchedulerClient)(nil).ReportPeerResult), varargs...)
}

// ReportPieceResult mocks base method.
func (m *MockSchedulerClient) ReportPieceResult(ctx context.Context, opts ...grpc.CallOption) (v1.Scheduler_ReportPieceResultClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReportPieceResult", varargs...)
	ret0, _ := ret[0].(v1.Scheduler_ReportPieceResultClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportPieceResult indicates an expected call of ReportPieceResult.
func (mr *MockSchedulerClientMockRecorder) ReportPieceResult(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportPieceResult", reflect.TypeOf((*MockSchedulerClient)(nil).ReportPieceResult), varargs...)
}

// StatTask mocks base method.
func (m *MockSchedulerClient) StatTask(ctx context.Context, in *v1.StatTaskRequest, opts ...grpc.CallOption) (*v1.Task, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatTask", varargs...)
	ret0, _ := ret[0].(*v1.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatTask indicates an expected call of StatTask.
func (mr *MockSchedulerClientMockRecorder) StatTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatTask", reflect.TypeOf((*MockSchedulerClient)(nil).StatTask), varargs...)
}

// MockScheduler_ReportPieceResultClient is a mock of Scheduler_ReportPieceResultClient interface.
type MockScheduler_ReportPieceResultClient struct {
	ctrl     *gomock.Controller
	recorder *MockScheduler_ReportPieceResultClientMockRecorder
}

// MockScheduler_ReportPieceResultClientMockRecorder is the mock recorder for MockScheduler_ReportPieceResultClient.
type MockScheduler_ReportPieceResultClientMockRecorder struct {
	mock *MockScheduler_ReportPieceResultClient
}

// NewMockScheduler_ReportPieceResultClient creates a new mock instance.
func NewMockScheduler_ReportPieceResultClient(ctrl *gomock.Controller) *MockScheduler_ReportPieceResultClient {
	mock := &MockScheduler_ReportPieceResultClient{ctrl: ctrl}
	mock.recorder = &MockScheduler_ReportPieceResultClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler_ReportPieceResultClient) EXPECT() *MockScheduler_ReportPieceResultClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockScheduler_ReportPieceResultClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockScheduler_ReportPieceResultClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).Context))
}

// Header mocks base method.
func (m *MockScheduler_ReportPieceResultClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockScheduler_ReportPieceResultClient) Recv() (*v1.PeerPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.PeerPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockScheduler_ReportPieceResultClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockScheduler_ReportPieceResultClient) Send(arg0 *v1.PieceResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockScheduler_ReportPieceResultClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockScheduler_ReportPieceResultClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockScheduler_ReportPieceResultClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockScheduler_ReportPieceResultClient)(nil).Trailer))
}

// MockSchedulerServer is a mock of SchedulerServer interface.
type MockSchedulerServer struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerServerMockRecorder
}

// MockSchedulerServerMockRecorder is the mock recorder for MockSchedulerServer.
type MockSchedulerServerMockRecorder struct {
	mock *MockSchedulerServer
}

// NewMockSchedulerServer creates a new mock instance.
func NewMockSchedulerServer(ctrl *gomock.Controller) *MockSchedulerServer {
	mock := &MockSchedulerServer{ctrl: ctrl}
	mock.recorder = &MockSchedulerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerServer) EXPECT() *MockSchedulerServerMockRecorder {
	return m.recorder
}

// AnnounceTask mocks base method.
func (m *MockSchedulerServer) AnnounceTask(arg0 context.Context, arg1 *v1.AnnounceTaskRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnnounceTask", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnnounceTask indicates an expected call of AnnounceTask.
func (mr *MockSchedulerServerMockRecorder) AnnounceTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnnounceTask", reflect.TypeOf((*MockSchedulerServer)(nil).AnnounceTask), arg0, arg1)
}

// LeaveTask mocks base method.
func (m *MockSchedulerServer) LeaveTask(arg0 context.Context, arg1 *v1.PeerTarget) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LeaveTask", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LeaveTask indicates an expected call of LeaveTask.
func (mr *MockSchedulerServerMockRecorder) LeaveTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LeaveTask", reflect.TypeOf((*MockSchedulerServer)(nil).LeaveTask), arg0, arg1)
}

// RegisterPeerTask mocks base method.
func (m *MockSchedulerServer) RegisterPeerTask(arg0 context.Context, arg1 *v1.PeerTaskRequest) (*v1.RegisterResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterPeerTask", arg0, arg1)
	ret0, _ := ret[0].(*v1.RegisterResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterPeerTask indicates an expected call of RegisterPeerTask.
func (mr *MockSchedulerServerMockRecorder) RegisterPeerTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPeerTask", reflect.TypeOf((*MockSchedulerServer)(nil).RegisterPeerTask), arg0, arg1)
}

// ReportPeerResult mocks base method.
func (m *MockSchedulerServer) ReportPeerResult(arg0 context.Context, arg1 *v1.PeerResult) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportPeerResult", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportPeerResult indicates an expected call of ReportPeerResult.
func (mr *MockSchedulerServerMockRecorder) ReportPeerResult(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportPeerResult", reflect.TypeOf((*MockSchedulerServer)(nil).ReportPeerResult), arg0, arg1)
}

// ReportPieceResult mocks base method.
func (m *MockSchedulerServer) ReportPieceResult(arg0 v1.Scheduler_ReportPieceResultServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportPieceResult", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReportPieceResult indicates an expected call of ReportPieceResult.
func (mr *MockSchedulerServerMockRecorder) ReportPieceResult(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportPieceResult", reflect.TypeOf((*MockSchedulerServer)(nil).ReportPieceResult), arg0)
}

// StatTask mocks base method.
func (m *MockSchedulerServer) StatTask(arg0 context.Context, arg1 *v1.StatTaskRequest) (*v1.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatTask", arg0, arg1)
	ret0, _ := ret[0].(*v1.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatTask indicates an expected call of StatTask.
func (mr *MockSchedulerServerMockRecorder) StatTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatTask", reflect.TypeOf((*MockSchedulerServer)(nil).StatTask), arg0, arg1)
}

// MockScheduler_ReportPieceResultServer is a mock of Scheduler_ReportPieceResultServer interface.
type MockScheduler_ReportPieceResultServer struct {
	ctrl     *gomock.Controller
	recorder *MockScheduler_ReportPieceResultServerMockRecorder
}

// MockScheduler_ReportPieceResultServerMockRecorder is the mock recorder for MockScheduler_ReportPieceResultServer.
type MockScheduler_ReportPieceResultServerMockRecorder struct {
	mock *MockScheduler_ReportPieceResultServer
}

// NewMockScheduler_ReportPieceResultServer creates a new mock instance.
func NewMockScheduler_ReportPieceResultServer(ctrl *gomock.Controller) *MockScheduler_ReportPieceResultServer {
	mock := &MockScheduler_ReportPieceResultServer{ctrl: ctrl}
	mock.recorder = &MockScheduler_ReportPieceResultServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScheduler_ReportPieceResultServer) EXPECT() *MockScheduler_ReportPieceResultServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockScheduler_ReportPieceResultServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockScheduler_ReportPieceResultServer) Recv() (*v1.PieceResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.PieceResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockScheduler_ReportPieceResultServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockScheduler_ReportPieceResultServer) Send(arg0 *v1.PeerPacket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockScheduler_ReportPieceResultServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockScheduler_ReportPieceResultServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockScheduler_ReportPieceResultServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockScheduler_ReportPieceResultServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockScheduler_ReportPieceResultServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockScheduler_ReportPieceResultServer)(nil).SetTrailer), arg0)
}
