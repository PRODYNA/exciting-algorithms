package tests

import (
	"excitingalgorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntensity(t *testing.T) {
	intensity, err := erlang.Intensity(200, 180, 60)
	// Intensity
	assert.Equal(t, intensity, float64(10), "Intensity are equal")
	assert.Equal(t, err, nil, "Error is nil")
	assert.NotEqual(t, intensity, 60, "Intensity not equal")

}
