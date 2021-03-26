package tests

import (
	"excitingalgorithm/erlang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfAgentsForSl(t *testing.T) {
	numberofAgentsForSl, err := erlang.erlang.NumberOfAgentsForSl(200, 180, 60, 60, 0.8)
	// NumberOfAgentsForSl
	assert.Equal(t, numberofAgentsForSl, 13, "NumberOfAgentsForSL are equal")
	assert.Equal(t, err, nil, "No Error in Number of Agents")
	assert.NotEqual(t, numberofAgentsForSl, 18, "NumberOfAgentsForSL not equal")

}
