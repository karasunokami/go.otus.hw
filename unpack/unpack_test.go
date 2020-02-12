package unpack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	result, err := Unpack("a4bc2d5e")
	assert.Nil(t, err)
	assert.Equal(t, "aaaabccddddde", result)

	result, err = Unpack("abcd")
	assert.Nil(t, err)
	assert.Equal(t, "abcd", result)

	result, err = Unpack("45")
	assert.EqualError(t, err, "invalid string: 45")
	assert.Equal(t, "", result)

	result, err = Unpack("7q")
	assert.EqualError(t, err, "invalid string: 7q")
	assert.Equal(t, "", result)

	result, err = Unpack("q3")
	assert.Nil(t, err)
	assert.Equal(t, "qqq", result)
}
