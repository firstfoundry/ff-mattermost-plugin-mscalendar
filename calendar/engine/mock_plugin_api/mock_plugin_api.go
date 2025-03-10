// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/firstfoundry/ff-mattermost-plugin-mscalendar/calendar/engine (interfaces: PluginAPI)

// Package mock_plugin_api is a generated GoMock package.
package mock_plugin_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v6/model"
)

// MockPluginAPI is a mock of PluginAPI interface.
type MockPluginAPI struct {
	ctrl     *gomock.Controller
	recorder *MockPluginAPIMockRecorder
}

// MockPluginAPIMockRecorder is the mock recorder for MockPluginAPI.
type MockPluginAPIMockRecorder struct {
	mock *MockPluginAPI
}

// NewMockPluginAPI creates a new mock instance.
func NewMockPluginAPI(ctrl *gomock.Controller) *MockPluginAPI {
	mock := &MockPluginAPI{ctrl: ctrl}
	mock.recorder = &MockPluginAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPluginAPI) EXPECT() *MockPluginAPIMockRecorder {
	return m.recorder
}

// CanLinkEventToChannel mocks base method.
func (m *MockPluginAPI) CanLinkEventToChannel(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanLinkEventToChannel", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanLinkEventToChannel indicates an expected call of CanLinkEventToChannel.
func (mr *MockPluginAPIMockRecorder) CanLinkEventToChannel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanLinkEventToChannel", reflect.TypeOf((*MockPluginAPI)(nil).CanLinkEventToChannel), arg0, arg1)
}

// GetMattermostUser mocks base method.
func (m *MockPluginAPI) GetMattermostUser(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMattermostUser", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMattermostUser indicates an expected call of GetMattermostUser.
func (mr *MockPluginAPIMockRecorder) GetMattermostUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMattermostUser", reflect.TypeOf((*MockPluginAPI)(nil).GetMattermostUser), arg0)
}

// GetMattermostUserByUsername mocks base method.
func (m *MockPluginAPI) GetMattermostUserByUsername(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMattermostUserByUsername", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMattermostUserByUsername indicates an expected call of GetMattermostUserByUsername.
func (mr *MockPluginAPIMockRecorder) GetMattermostUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMattermostUserByUsername", reflect.TypeOf((*MockPluginAPI)(nil).GetMattermostUserByUsername), arg0)
}

// GetMattermostUserStatus mocks base method.
func (m *MockPluginAPI) GetMattermostUserStatus(arg0 string) (*model.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMattermostUserStatus", arg0)
	ret0, _ := ret[0].(*model.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMattermostUserStatus indicates an expected call of GetMattermostUserStatus.
func (mr *MockPluginAPIMockRecorder) GetMattermostUserStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMattermostUserStatus", reflect.TypeOf((*MockPluginAPI)(nil).GetMattermostUserStatus), arg0)
}

// GetMattermostUserStatusesByIds mocks base method.
func (m *MockPluginAPI) GetMattermostUserStatusesByIds(arg0 []string) ([]*model.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMattermostUserStatusesByIds", arg0)
	ret0, _ := ret[0].([]*model.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMattermostUserStatusesByIds indicates an expected call of GetMattermostUserStatusesByIds.
func (mr *MockPluginAPIMockRecorder) GetMattermostUserStatusesByIds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMattermostUserStatusesByIds", reflect.TypeOf((*MockPluginAPI)(nil).GetMattermostUserStatusesByIds), arg0)
}

// GetMattermostUserTeams mocks base method.
func (m *MockPluginAPI) GetMattermostUserTeams(arg0 string) ([]*model.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMattermostUserTeams", arg0)
	ret0, _ := ret[0].([]*model.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMattermostUserTeams indicates an expected call of GetMattermostUserTeams.
func (mr *MockPluginAPIMockRecorder) GetMattermostUserTeams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMattermostUserTeams", reflect.TypeOf((*MockPluginAPI)(nil).GetMattermostUserTeams), arg0)
}

// GetPost mocks base method.
func (m *MockPluginAPI) GetPost(arg0 string) (*model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", arg0)
	ret0, _ := ret[0].(*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockPluginAPIMockRecorder) GetPost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockPluginAPI)(nil).GetPost), arg0)
}

// IsSysAdmin mocks base method.
func (m *MockPluginAPI) IsSysAdmin(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSysAdmin", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSysAdmin indicates an expected call of IsSysAdmin.
func (mr *MockPluginAPIMockRecorder) IsSysAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSysAdmin", reflect.TypeOf((*MockPluginAPI)(nil).IsSysAdmin), arg0)
}

// PublishWebsocketEvent mocks base method.
func (m *MockPluginAPI) PublishWebsocketEvent(arg0, arg1 string, arg2 map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishWebsocketEvent", arg0, arg1, arg2)
}

// PublishWebsocketEvent indicates an expected call of PublishWebsocketEvent.
func (mr *MockPluginAPIMockRecorder) PublishWebsocketEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWebsocketEvent", reflect.TypeOf((*MockPluginAPI)(nil).PublishWebsocketEvent), arg0, arg1, arg2)
}

// SearchLinkableChannelForUser mocks base method.
func (m *MockPluginAPI) SearchLinkableChannelForUser(arg0, arg1, arg2 string) ([]*model.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchLinkableChannelForUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchLinkableChannelForUser indicates an expected call of SearchLinkableChannelForUser.
func (mr *MockPluginAPIMockRecorder) SearchLinkableChannelForUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchLinkableChannelForUser", reflect.TypeOf((*MockPluginAPI)(nil).SearchLinkableChannelForUser), arg0, arg1, arg2)
}

// UpdateMattermostUserStatus mocks base method.
func (m *MockPluginAPI) UpdateMattermostUserStatus(arg0, arg1 string) (*model.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMattermostUserStatus", arg0, arg1)
	ret0, _ := ret[0].(*model.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMattermostUserStatus indicates an expected call of UpdateMattermostUserStatus.
func (mr *MockPluginAPIMockRecorder) UpdateMattermostUserStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMattermostUserStatus", reflect.TypeOf((*MockPluginAPI)(nil).UpdateMattermostUserStatus), arg0, arg1)
}
