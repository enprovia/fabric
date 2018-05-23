// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Capabilities is an autogenerated mock type for the Capabilities type
type Capabilities struct {
	mock.Mock
}

// CollectionUpgrade provides a mock function with given fields:
func (_m *Capabilities) CollectionUpgrade() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ForbidDuplicateTXIdInBlock provides a mock function with given fields:
func (_m *Capabilities) ForbidDuplicateTXIdInBlock() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// KeyLevelEndorsement provides a mock function with given fields:
func (_m *Capabilities) KeyLevelEndorsement() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MetadataLifecycle provides a mock function with given fields:
func (_m *Capabilities) MetadataLifecycle() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PrivateChannelData provides a mock function with given fields:
func (_m *Capabilities) PrivateChannelData() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ResourcesTree provides a mock function with given fields:
func (_m *Capabilities) ResourcesTree() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Supported provides a mock function with given fields:
func (_m *Capabilities) Supported() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// V1_1Validation provides a mock function with given fields:
func (_m *Capabilities) V1_1Validation() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// V1_2Validation provides a mock function with given fields:
func (_m *Capabilities) V1_2Validation() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
