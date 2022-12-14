// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// PortRepoService is an autogenerated mock type for the PortRepoService type
type PortRepoService struct {
	mock.Mock
}

// DeletePortByID provides a mock function with given fields: ctx, id
func (_m *PortRepoService) DeletePortByID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPortByID provides a mock function with given fields: ctx, id
func (_m *PortRepoService) GetPortByID(ctx context.Context, id string) (models.PortDetails, error) {
	ret := _m.Called(ctx, id)

	var r0 models.PortDetails
	if rf, ok := ret.Get(0).(func(context.Context, string) models.PortDetails); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.PortDetails)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPorts provides a mock function with given fields: ctx
func (_m *PortRepoService) GetPorts(ctx context.Context) ([]models.PortDetails, error) {
	ret := _m.Called(ctx)

	var r0 []models.PortDetails
	if rf, ok := ret.Get(0).(func(context.Context) []models.PortDetails); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.PortDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertPorts provides a mock function with given fields: ctx, portsDetails
func (_m *PortRepoService) InsertPorts(ctx context.Context, portsDetails []models.PortDetails) error {
	ret := _m.Called(ctx, portsDetails)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []models.PortDetails) error); ok {
		r0 = rf(ctx, portsDetails)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePortByID provides a mock function with given fields: ctx, id, port
func (_m *PortRepoService) UpdatePortByID(ctx context.Context, id string, port *models.PortDetails) error {
	ret := _m.Called(ctx, id, port)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *models.PortDetails) error); ok {
		r0 = rf(ctx, id, port)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPortRepoService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPortRepoService creates a new instance of PortRepoService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPortRepoService(t mockConstructorTestingTNewPortRepoService) *PortRepoService {
	mock := &PortRepoService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
