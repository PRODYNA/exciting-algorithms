package tests

import (
	"excitingalgorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfAgentsForAsa(t *testing.T) {
	numberOfAgentsForAsa, err := erlang.NumberOfAgentsForAsa(200, 180, 60, 120)
	// NumberOfAgentsForAsa
	assert.Equal(t, numberOfAgentsForAsa, 12, "NumberOfAgentsForAsa are equal")
	assert.Equal(t, err, nil, "NumberOfAgentsForAsa are equal")
	assert.NotEqual(t, numberOfAgentsForAsa, 19, "NumberOfAgentsForAsa not equal")

}
