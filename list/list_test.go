package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList(t *testing.T) {
	list := new(List)

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)

	assert.Equal(t, list.Len(), 3)
	assert.Equal(t, list.Last().Value, 1)
	assert.Equal(t, list.First().Value, 3)
}
