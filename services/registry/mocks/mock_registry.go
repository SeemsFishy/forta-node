// Code generated by MockGen. DO NOT EDIT.
// Source: services/registry/registry.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	context "context"
	domain "forta-network/forta-node/domain"
	regtypes "forta-network/forta-node/services/registry/regtypes"
	big "math/big"
	reflect "reflect"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
)

// MockAgentRegistryCaller is a mock of AgentRegistryCaller interface.
type MockAgentRegistryCaller struct {
	ctrl     *gomock.Controller
	recorder *MockAgentRegistryCallerMockRecorder
}

// MockAgentRegistryCallerMockRecorder is the mock recorder for MockAgentRegistryCaller.
type MockAgentRegistryCallerMockRecorder struct {
	mock *MockAgentRegistryCaller
}

// NewMockAgentRegistryCaller creates a new mock instance.
func NewMockAgentRegistryCaller(ctrl *gomock.Controller) *MockAgentRegistryCaller {
	mock := &MockAgentRegistryCaller{ctrl: ctrl}
	mock.recorder = &MockAgentRegistryCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentRegistryCaller) EXPECT() *MockAgentRegistryCallerMockRecorder {
	return m.recorder
}

// AgentLatestVersion mocks base method.
func (m *MockAgentRegistryCaller) AgentLatestVersion(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentLatestVersion", opts, arg0)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentLatestVersion indicates an expected call of AgentLatestVersion.
func (mr *MockAgentRegistryCallerMockRecorder) AgentLatestVersion(opts, arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentLatestVersion", reflect.TypeOf((*MockAgentRegistryCaller)(nil).AgentLatestVersion), opts, arg0)
}

// AgentReference mocks base method.
func (m *MockAgentRegistryCaller) AgentReference(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentReference", opts, arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentReference indicates an expected call of AgentReference.
func (mr *MockAgentRegistryCallerMockRecorder) AgentReference(opts, arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentReference", reflect.TypeOf((*MockAgentRegistryCaller)(nil).AgentReference), opts, arg0, arg1)
}

// MockScannerRegistryCaller is a mock of ScannerRegistryCaller interface.
type MockScannerRegistryCaller struct {
	ctrl     *gomock.Controller
	recorder *MockScannerRegistryCallerMockRecorder
}

// MockScannerRegistryCallerMockRecorder is the mock recorder for MockScannerRegistryCaller.
type MockScannerRegistryCallerMockRecorder struct {
	mock *MockScannerRegistryCaller
}

// NewMockScannerRegistryCaller creates a new mock instance.
func NewMockScannerRegistryCaller(ctrl *gomock.Controller) *MockScannerRegistryCaller {
	mock := &MockScannerRegistryCaller{ctrl: ctrl}
	mock.recorder = &MockScannerRegistryCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScannerRegistryCaller) EXPECT() *MockScannerRegistryCallerMockRecorder {
	return m.recorder
}

// AgentAt mocks base method.
func (m *MockScannerRegistryCaller) AgentAt(opts *bind.CallOpts, scanner common.Address, index *big.Int) ([32]byte, *big.Int, bool, string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentAt", opts, scanner, index)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(*big.Int)
	ret2, _ := ret[2].(bool)
	ret3, _ := ret[3].(string)
	ret4, _ := ret[4].(bool)
	ret5, _ := ret[5].(error)
	return ret0, ret1, ret2, ret3, ret4, ret5
}

// AgentAt indicates an expected call of AgentAt.
func (mr *MockScannerRegistryCallerMockRecorder) AgentAt(opts, scanner, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentAt", reflect.TypeOf((*MockScannerRegistryCaller)(nil).AgentAt), opts, scanner, index)
}

// AgentLength mocks base method.
func (m *MockScannerRegistryCaller) AgentLength(opts *bind.CallOpts, scanner common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentLength", opts, scanner)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentLength indicates an expected call of AgentLength.
func (mr *MockScannerRegistryCallerMockRecorder) AgentLength(opts, scanner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentLength", reflect.TypeOf((*MockScannerRegistryCaller)(nil).AgentLength), opts, scanner)
}

// GetAgentListHash mocks base method.
func (m *MockScannerRegistryCaller) GetAgentListHash(opts *bind.CallOpts, scanner common.Address) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentListHash", opts, scanner)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentListHash indicates an expected call of GetAgentListHash.
func (mr *MockScannerRegistryCallerMockRecorder) GetAgentListHash(opts, scanner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentListHash", reflect.TypeOf((*MockScannerRegistryCaller)(nil).GetAgentListHash), opts, scanner)
}

// MockIPFSClient is a mock of IPFSClient interface.
type MockIPFSClient struct {
	ctrl     *gomock.Controller
	recorder *MockIPFSClientMockRecorder
}

// MockIPFSClientMockRecorder is the mock recorder for MockIPFSClient.
type MockIPFSClientMockRecorder struct {
	mock *MockIPFSClient
}

// NewMockIPFSClient creates a new mock instance.
func NewMockIPFSClient(ctrl *gomock.Controller) *MockIPFSClient {
	mock := &MockIPFSClient{ctrl: ctrl}
	mock.recorder = &MockIPFSClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPFSClient) EXPECT() *MockIPFSClientMockRecorder {
	return m.recorder
}

// GetAgentFile mocks base method.
func (m *MockIPFSClient) GetAgentFile(cid string) (*regtypes.AgentFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentFile", cid)
	ret0, _ := ret[0].(*regtypes.AgentFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAgentFile indicates an expected call of GetAgentFile.
func (mr *MockIPFSClientMockRecorder) GetAgentFile(cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentFile", reflect.TypeOf((*MockIPFSClient)(nil).GetAgentFile), cid)
}

// MockEthClient is a mock of EthClient interface.
type MockEthClient struct {
	ctrl     *gomock.Controller
	recorder *MockEthClientMockRecorder
}

// MockEthClientMockRecorder is the mock recorder for MockEthClient.
type MockEthClientMockRecorder struct {
	mock *MockEthClient
}

// NewMockEthClient creates a new mock instance.
func NewMockEthClient(ctrl *gomock.Controller) *MockEthClient {
	mock := &MockEthClient{ctrl: ctrl}
	mock.recorder = &MockEthClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEthClient) EXPECT() *MockEthClientMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method.
func (m *MockEthClient) BlockByHash(ctx context.Context, hash string) (*domain.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash)
	ret0, _ := ret[0].(*domain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockEthClientMockRecorder) BlockByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockEthClient)(nil).BlockByHash), ctx, hash)
}

// BlockByNumber mocks base method.
func (m *MockEthClient) BlockByNumber(ctx context.Context, number *big.Int) (*domain.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number)
	ret0, _ := ret[0].(*domain.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockEthClientMockRecorder) BlockByNumber(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockEthClient)(nil).BlockByNumber), ctx, number)
}

// BlockNumber mocks base method.
func (m *MockEthClient) BlockNumber(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockEthClientMockRecorder) BlockNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockEthClient)(nil).BlockNumber), ctx)
}

// ChainID mocks base method.
func (m *MockEthClient) ChainID(ctx context.Context) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID", ctx)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainID indicates an expected call of ChainID.
func (mr *MockEthClientMockRecorder) ChainID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockEthClient)(nil).ChainID), ctx)
}

// Close mocks base method.
func (m *MockEthClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockEthClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEthClient)(nil).Close))
}

// GetLogs mocks base method.
func (m *MockEthClient) GetLogs(ctx context.Context, hash string) ([]domain.LogEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", ctx, hash)
	ret0, _ := ret[0].([]domain.LogEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockEthClientMockRecorder) GetLogs(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockEthClient)(nil).GetLogs), ctx, hash)
}

// TraceBlock mocks base method.
func (m *MockEthClient) TraceBlock(ctx context.Context, number *big.Int) ([]domain.Trace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceBlock", ctx, number)
	ret0, _ := ret[0].([]domain.Trace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TraceBlock indicates an expected call of TraceBlock.
func (mr *MockEthClientMockRecorder) TraceBlock(ctx, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBlock", reflect.TypeOf((*MockEthClient)(nil).TraceBlock), ctx, number)
}

// TransactionReceipt mocks base method.
func (m *MockEthClient) TransactionReceipt(ctx context.Context, txHash string) (*domain.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*domain.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockEthClientMockRecorder) TransactionReceipt(ctx, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockEthClient)(nil).TransactionReceipt), ctx, txHash)
}
