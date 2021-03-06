package tests

import (
	"github.com/prodyna/exciting-algorithms/erlang"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestOccupancy(t *testing.T) {

	// Occupancy
	assert.Equal(t, math.Round(erlang.Occupancy(23, 20)/0.001)*0.001, 0.870, "Occupancy are equal")
	assert.NotEqual(t, erlang.Occupancy(23, 20), 0.869565, "Occupancy not equal")

}
