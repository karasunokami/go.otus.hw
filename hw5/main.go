package hw5

import (
	"errors"
	"math"
	"sync"
)

func Run(tasks []task, goroutinesCount int, errorsLimit int) error {
	if goroutinesCount <= 0 {
		return errors.New("cant handle non or negative goroutines count")
	}

	wg := &sync.WaitGroup{}

	errorsChan := make(chan error)
	tasksChan := make(chan task)
	successChan := make(chan struct{})

	errorsCount := 0
	chunks := arrayChunk(tasks, goroutinesCount)

	wg.Add(len(tasks))
	for i := 0; i < goroutinesCount; i++ {
		go worker(tasksChan, errorsChan, successChan, wg)
	}

	for _, chunk := range chunks {
		for i := 0; i < len(chunk); i++ {
			tasksChan <- chunk[i]
		}

		for i := 0; i < len(chunk); i++ {
			select {
				case <-errorsChan:
					errorsCount++
					if errorsCount > errorsLimit {
						return errors.New("errors limit exceeded")
					}

				case <-successChan:
			}
		}
	}

	close(tasksChan)
	wg.Wait()

	return nil
}

func arrayChunk(arr []task, size int) [][]task {
	if size < 1 {
		panic("size: cannot be less than 1")
	}

	length := len(arr)
	chunks := int(math.Ceil(float64(length) / float64(size)))

	var chunk [][]task
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

func worker(tasksChan <- chan task, errorsChan chan <- error, successChan chan <- struct{}, wg *sync.WaitGroup) {
	for t := range tasksChan {
		err := t()

		if err != nil {
			errorsChan <- err
		} else {
			successChan <- struct{}{}
		}
		wg.Done()
	}
}