package function_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	result := sum(3, 4)
	assert.Equal(t, 7, result)
}

func TestName1(t *testing.T) {
	result := sum(3, 4)
	assert.NotEqual(t, 8, result)
}
