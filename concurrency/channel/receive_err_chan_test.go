package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiveErrFromChan(t *testing.T) {
	result := ReceiveErrFromChan(5, true)
	assert.Equal(t, "n is even", result)

	result = ReceiveErrFromChan(3, false)
	assert.Equal(t, "", result)
}
