package tests

import (
	"github.com/prodyna/exciting-algorithms/erlang"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestErlangc(t *testing.T) {

	// ErlangC
	assert.Equal(t, math.Round(erlang.ErlangC(25, 20)/0.001)*0.001, 0.261, "ErlangC are equal")
	assert.NotEqual(t, erlang.ErlangC(25, 20), 0.261379, "ErlangC not equal")

}
