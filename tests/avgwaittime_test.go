package tests

import (
	erlang2 "excitingalgorithm/erlang"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAvgWaitTime(t *testing.T) {

	avgWaitTime, err := erlang2.AvgWaitTime(25, 200, 180, 30)
	// AvgWaitTime
	assert.Equal(t, math.Round(avgWaitTime/0.001)*0.001, 9.410, "AvgWaitTime are equal")
	assert.Equal(t, err, nil, "Error is nil")
	assert.NotEqual(t, math.Round(avgWaitTime/0.001)*0.001, 10, "AvgWaitTime not equal")

}
