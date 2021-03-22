package tests

import (
	"excitingalgorithm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfAgentsForAsa(t *testing.T) {

	// NumberOfAgentsForAsa
	assert.Equal(t, erlang.NumberOfAgentsForAsa(200, 180, 60, 120), 12, "NumberOfAgentsForAsa are equal")
	assert.NotEqual(t, erlang.NumberOfAgentsForAsa(200, 180, 60, 120), 19, "NumberOfAgentsForAsa not equal")

}
