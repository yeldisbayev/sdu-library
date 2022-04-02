// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/yeldisbayev/thrid-group/library/internal/domain"
)

// BookRepo is an autogenerated mock type for the BookRepo type
type BookRepo struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, id
func (_m *BookRepo) Get(ctx context.Context, id string) (*domain.Book, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Book
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Book); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: ctx, id
func (_m *BookRepo) Remove(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store provides a mock function with given fields: ctx, book
func (_m *BookRepo) Store(ctx context.Context, book domain.Book) (*domain.Book, error) {
	ret := _m.Called(ctx, book)

	var r0 *domain.Book
	if rf, ok := ret.Get(0).(func(context.Context, domain.Book) *domain.Book); ok {
		r0 = rf(ctx, book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.Book) error); ok {
		r1 = rf(ctx, book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}