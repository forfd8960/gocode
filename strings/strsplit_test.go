package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrSplit_Split(t *testing.T) {
	ss := NewStrSplit("a b c", " ")

	iter := ss.Split()

	assert.Equal(t, "a", iter.Next())
	assert.Equal(t, "b", iter.Next())
	assert.Equal(t, "c", iter.Next())
	assert.Equal(t, "", iter.Next())
}
