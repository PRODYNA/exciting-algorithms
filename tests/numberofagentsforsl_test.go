package tests

import (
	"excitingalgorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfAgentsForSl(t *testing.T) {
	//// ServiceLevel
	//assert.Equal(t, ServiceLevel(25, 200,180,60), -4002690070144283688661186947383997812612770114231070593941097979595071905906726797312.000000, "ServiceLevel are equal")
	//assert.NotEqual(t, ServiceLevel(25, 200,180,60), -4002690070144283688661186947383997812612770114231070593941097979595071905906726797312.000000, "SeviceLevel not equal")

	// NumberOfAgentsForSl
	assert.Equal(t, erlang.NumberOfAgentsForSl(200, 180, 60, 60, 0.8), 13, "NumberOfAgentsForSL are equal")
	assert.NotEqual(t, erlang.NumberOfAgentsForSl(200, 180, 60, 60, 0.8), 18, "NumberOfAgentsForSL not equal")

}
