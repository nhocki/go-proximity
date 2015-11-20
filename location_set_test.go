package proximity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const key = "go-proximity:test-set"

func TestSuccessfulAdd(t *testing.T) {
	var (
		lat          = 123.0
		lng          = -123.0
		locationName = "Test Location"
	)

	client := &testClient{}
	client.On("ZAdd", key, encode(lat, lng), locationName).Return(1, nil)

	set := &LocationSet{
		Name:   key,
		client: client,
	}

	err := set.Add(locationName, lat, lng)
	client.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestUnsuccessfulAdd(t *testing.T) {
	var (
		lat          = 123.0
		lng          = -123.0
		locationName = "Test Location"
	)

	client := &testClient{}
	err := errors.New("Some Redis Error")
	client.On("ZAdd", key, encode(lat, lng), locationName).Return(0, err)

	set := &LocationSet{
		Name:   key,
		client: client,
	}

	err = set.Add(locationName, lat, lng)
	client.AssertExpectations(t)
	assert.NotNil(t, err)
}
