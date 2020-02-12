package ntptime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrint(t *testing.T) {
	_, _, err := Get()

	assert.Equal(t, err, nil)
}
