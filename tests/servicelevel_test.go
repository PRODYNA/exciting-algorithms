package tests

import (
	"github.com/prodyna/exciting-algorithms/erlang"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestServiceLevel(t *testing.T) {
	// ServiceLevel
	serviceLevel, err := erlang.ServiceLevel(25, 200, 180, 60, 60)

	assert.Equal(t, math.Round(serviceLevel/0.0000001)*0.0000001, 0.9999992, "ServiceLevel are equal")
	assert.Equal(t, err, nil, "ServiceLevel are equal")
	assert.NotEqual(t, serviceLevel, 0.999, "SeviceLevel not equal")
}
