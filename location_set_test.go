package proximity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	set          *LocationSet
	lat          = 123.0
	lng          = -123.0
	radius       = 5000.0
	lowerBound   = -1.0
	upperBound   = 1.0
	locationName = "Test Location"
)

func init() {
	set = &LocationSet{
		Name: key,
		IntervalFinder: func(lat, lng, radius float64) []Int64arr {
			return []Int64arr{
				Int64arr{int64(lowerBound), int64(upperBound)},
			}
		},
	}
}

func TestSuccessfulAdd(t *testing.T) {
	client := &testClient{}
	client.On("ZAdd", key, encode(lat, lng), locationName).Return(1, nil)
	set.client = client

	err := set.Add(locationName, lat, lng)

	client.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestUnsuccessfulAdd(t *testing.T) {
	err := errors.New("Some Redis Error")

	client := &testClient{}
	client.On("ZAdd", key, encode(lat, lng), locationName).Return(0, err)
	set.client = client

	err = set.Add(locationName, lat, lng)

	client.AssertExpectations(t)
	assert.NotNil(t, err)
}

func TestSuccessfulNear(t *testing.T) {
	client := &testClient{}
	client.On("ZRangeByScore", set.Name, lowerBound, upperBound).Return([]string{"Some Value"}, nil)
	set.client = client

	results, err := set.Near(lat, lng, radius)
	client.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, "Some Value", results[0])
}

func TestUnsuccessfulNear(t *testing.T) {
	err := errors.New("Some Redis Error")

	client := &testClient{}
	client.On("ZRangeByScore", set.Name, lowerBound, upperBound).Return([]string{}, err)
	set.client = client

	results, err := set.Near(lat, lng, radius)
	client.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, 0, len(results))
}
