package tests

import (
	"excitingalgorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntensity(t *testing.T) {
	// Intensity
	assert.Equal(t, erlang.Intensity(200, 180, 60), float64(10), "Intensity are equal")
	assert.NotEqual(t, erlang.Intensity(200, 180, 60), 60, "Intensity not equal")

}
