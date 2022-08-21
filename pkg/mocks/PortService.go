// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	models "github.com/BoggalaPrabhakar007/golang-assignment/pkg/models"
)

// PortService is an autogenerated mock type for the PortService type
type PortService struct {
	mock.Mock
}

// DeletePortByID provides a mock function with given fields: ctx, id
func (_m *PortService) DeletePortByID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPortDataByID provides a mock function with given fields: ctx, id
func (_m *PortService) GetPortDataByID(ctx context.Context, id string) (models.PortDetails, error) {
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

// GetPortsData provides a mock function with given fields: ctx
func (_m *PortService) GetPortsData(ctx context.Context) ([]models.PortDetails, error) {
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

// InsertPortData provides a mock function with given fields: ctx, r
func (_m *PortService) InsertPortData(ctx context.Context, r *http.Request) error {
	ret := _m.Called(ctx, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) error); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePortByID provides a mock function with given fields: ctx, port, id
func (_m *PortService) UpdatePortByID(ctx context.Context, port models.PortDetails, id string) error {
	ret := _m.Called(ctx, port, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.PortDetails, string) error); ok {
		r0 = rf(ctx, port, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPortService interface {
	mock.TestingT
	Cleanup(func())
}

// NewPortService creates a new instance of PortService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPortService(t mockConstructorTestingTNewPortService) *PortService {
	mock := &PortService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
