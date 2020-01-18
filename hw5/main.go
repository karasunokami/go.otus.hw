package hw5

import (
	"errors"
	"math"
)

func Run(tasks []func() error, goroutinesCount int, errorsLimit int) error {
	errorsCount := 0
	errorsChan := make(chan error)
	chunk := arrayChunk(tasks, goroutinesCount)

	for _, tasks := range chunk {
		for _, task := range tasks {
			go func(task func() error) {
				err := task()
				errorsChan <- err
			}(task)
		}

		for i := 0; i < len(tasks); i++ {
			err := <-errorsChan

			if err != nil {
				errorsCount++
			}

			if errorsCount == errorsLimit {
				return errors.New("errors limit exceeded")
			}
		}
	}

	return nil
}

func arrayChunk(arr []func() error, size int) [][]func() error {
	if size < 1 {
		panic("size: cannot be less than 1")
	}

	length := len(arr)
	chunks := int(math.Ceil(float64(length) / float64(size)))

	var chunk [][]func() error
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		chunk = append(chunk, arr[i * size:end])
		i++
	}

	return chunk
}