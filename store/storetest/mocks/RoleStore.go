// Code generated by mockery v1.0.0

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// RoleStore is an autogenerated mock type for the RoleStore type
type RoleStore struct {
	mock.Mock
}

// Get provides a mock function with given fields: roleId
func (_m *RoleStore) Get(roleId string) store.StoreChannel {
	ret := _m.Called(roleId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(roleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByName provides a mock function with given fields: name
func (_m *RoleStore) GetByName(name string) store.StoreChannel {
	ret := _m.Called(name)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetByNames provides a mock function with given fields: names
func (_m *RoleStore) GetByNames(names []string) store.StoreChannel {
	ret := _m.Called(names)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func([]string) store.StoreChannel); ok {
		r0 = rf(names)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// PermanentDeleteAll provides a mock function with given fields:
func (_m *RoleStore) PermanentDeleteAll() store.StoreChannel {
	ret := _m.Called()

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func() store.StoreChannel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Save provides a mock function with given fields: role
func (_m *RoleStore) Save(role *model.Role) store.StoreChannel {
	ret := _m.Called(role)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Role) store.StoreChannel); ok {
		r0 = rf(role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
