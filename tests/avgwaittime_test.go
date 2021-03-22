package tests

import (
	"excitingalgorithm"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAvgWaitTime(t *testing.T) {

	// AvgWaitTime
	assert.Equal(t, math.Round(erlang.AvgWaitTime(25, 200, 180, 30)/0.001)*0.001, 9.410, "AvgWaitTime are equal")
	assert.NotEqual(t, erlang.AvgWaitTime(25, 200, 180, 30), 10, "AvgWaitTime not equal")

}
