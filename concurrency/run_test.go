package concurrency

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

	tasks := []task{task1, task2, task3, task4}

	err := Run(tasks, 3, 3)
	assert.Equal(t, nil, err)
}

func TestRunWithError(t *testing.T) {
	task1 := func() error { return nil }
	task2 := func() error { return errors.New("error") }
	task3 := func() error { return errors.New("error") }
	task4 := func() error { return errors.New("error") }

	tasks := []task{task1, task2, task3, task4}

	err := Run(tasks, 3, 2)
	expected := errors.New("errors limit exceeded")
	assert.Equal(t, expected, err)
}

func TestRunWithZeroCountOfGoroutines(t *testing.T) {
	var tasks []task

	err := Run(tasks, 0, 2)
	expected := errors.New("cant handle non or negative goroutines count")
	assert.Equal(t, expected, err)
}
