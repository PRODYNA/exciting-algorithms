package tests

import (
	"github.com/prodyna/exciting-algorithms/erlang"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestErlangb(t *testing.T) {

	// ErlangB
	assert.Equal(t, math.Round(erlang.ErlangB(200, 180)/0.001)*0.001, 0.012, "ErlangB are equal")
	assert.NotEqual(t, erlang.ErlangB(200, 180), 0.011592, "ErlangB not equal")

}
