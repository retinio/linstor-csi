// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/LINBIT/golinstor/client"

	mock "github.com/stretchr/testify/mock"
)

// BackupProvider is an autogenerated mock type for the BackupProvider type
type BackupProvider struct {
	mock.Mock
}

// Abort provides a mock function with given fields: ctx, remoteName, request
func (_m *BackupProvider) Abort(ctx context.Context, remoteName string, request client.BackupAbortRequest) error {
	ret := _m.Called(ctx, remoteName, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, client.BackupAbortRequest) error); ok {
		r0 = rf(ctx, remoteName, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: ctx, remoteName, request
func (_m *BackupProvider) Create(ctx context.Context, remoteName string, request client.BackupCreate) (string, error) {
	ret := _m.Called(ctx, remoteName, request)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, client.BackupCreate) string); ok {
		r0 = rf(ctx, remoteName, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, client.BackupCreate) error); ok {
		r1 = rf(ctx, remoteName, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAll provides a mock function with given fields: ctx, remoteName, filter
func (_m *BackupProvider) DeleteAll(ctx context.Context, remoteName string, filter client.BackupDeleteOpts) error {
	ret := _m.Called(ctx, remoteName, filter)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, client.BackupDeleteOpts) error); ok {
		r0 = rf(ctx, remoteName, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, remoteName, rscName, snapName
func (_m *BackupProvider) GetAll(ctx context.Context, remoteName string, rscName string, snapName string) (*client.BackupList, error) {
	ret := _m.Called(ctx, remoteName, rscName, snapName)

	var r0 *client.BackupList
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *client.BackupList); ok {
		r0 = rf(ctx, remoteName, rscName, snapName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.BackupList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, remoteName, rscName, snapName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Info provides a mock function with given fields: ctx, remoteName, request
func (_m *BackupProvider) Info(ctx context.Context, remoteName string, request client.BackupInfoRequest) (*client.BackupInfo, error) {
	ret := _m.Called(ctx, remoteName, request)

	var r0 *client.BackupInfo
	if rf, ok := ret.Get(0).(func(context.Context, string, client.BackupInfoRequest) *client.BackupInfo); ok {
		r0 = rf(ctx, remoteName, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.BackupInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, client.BackupInfoRequest) error); ok {
		r1 = rf(ctx, remoteName, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Restore provides a mock function with given fields: ctx, remoteName, request
func (_m *BackupProvider) Restore(ctx context.Context, remoteName string, request client.BackupRestoreRequest) error {
	ret := _m.Called(ctx, remoteName, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, client.BackupRestoreRequest) error); ok {
		r0 = rf(ctx, remoteName, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ship provides a mock function with given fields: ctx, remoteName, request
func (_m *BackupProvider) Ship(ctx context.Context, remoteName string, request client.BackupShipRequest) error {
	ret := _m.Called(ctx, remoteName, request)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, client.BackupShipRequest) error); ok {
		r0 = rf(ctx, remoteName, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
