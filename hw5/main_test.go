package hw5

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunWithNoErrors(t *testing.T) {
	task1 := func() error { return nil }
	task2 := func() error { return errors.New("error") }
	task3 := func() error { return nil }
	task4 := func() error { return errors.New("error") }

	tasks := []func() error{task1, task2, task3, task4}

	err := Run(tasks, 3, 3)
	assert.Equal(t, nil, err)
}

func TestRunWithError(t *testing.T) {
	task1 := func() error { return nil }
	task2 := func() error { return errors.New("error") }
	task3 := func() error { return nil }
	task4 := func() error { return errors.New("error") }

	tasks := []func() error{task1, task2, task3, task4}

	err := Run(tasks, 3, 2)
	expected := errors.New("errors limit exceeded")
	assert.Equal(t, expected, err)
}
