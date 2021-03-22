package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErlangAlgorithm(t *testing.T) {
	// Intensity
	assert.Equal(t, Intensity(200, 180), 600.000000, "Intensity are equal")
	assert.NotEqual(t, Intensity(200, 180), 600.000000, "Intensity not equal")

	// Occupancy
	assert.Equal(t, Occupancy(23, 20), 0.869565, "Occupancy are equal")
	assert.NotEqual(t, Occupancy(23, 20), 0.869565, "Occupancy not equal")

	// ErlangB
	assert.Equal(t, ErlangB(200, 180), 0.011592, "ErlangB are equal")
	assert.NotEqual(t, ErlangB(200, 180), 0.011592, "ErlangB not equal")

	// ErlangC
	assert.Equal(t, ErlangC(25, 20), 0.261379, "ErlangC are equal")
	assert.NotEqual(t, ErlangC(25, 20), 0.261379, "ErlangC not equal")

	// AvgWaitTime
	assert.Equal(t, AvgWaitTime(25, 200, 180), -7.214080, "AvgWaitTime are equal")
	assert.NotEqual(t, AvgWaitTime(25, 200, 180), -7.214080, "AvgWaitTime not equal")

	//// ServiceLevel
	//assert.Equal(t, ServiceLevel(25, 200,180,60), -4002690070144283688661186947383997812612770114231070593941097979595071905906726797312.000000, "ServiceLevel are equal")
	//assert.NotEqual(t, ServiceLevel(25, 200,180,60), -4002690070144283688661186947383997812612770114231070593941097979595071905906726797312.000000, "SeviceLevel not equal")

	// NumberOfAgentsForSl
	assert.Equal(t, NumberOfAgentsForSl(200, 180, 60, 0.8), 605, "NumberOfAgentsForSL are equal")
	assert.NotEqual(t, NumberOfAgentsForSl(200, 180, 60, 0.8), "NumberOfAgentsForSL not equal")

	// NumberOfAgentsForAsa
	assert.Equal(t, NumberOfAgentsForAsa(200, 180, 120), 602, "NumberOfAgentsForAsa are equal")
	assert.NotEqual(t, NumberOfAgentsForAsa(200, 180, 120), 602, "NumberOfAgentsForAsa not equal")

}
