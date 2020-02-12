package concurrency

import (
	"errors"
	"sync"
)

func Run(tasks []task, goroutinesCount int, errorsLimit int) error {
	if goroutinesCount <= 0 {
		return errors.New("cant handle non or negative goroutines count")
	}

	wg := &sync.WaitGroup{}

	errorsChan := make(chan error)
	tasksChan := make(chan task, goroutinesCount)
	successChan := make(chan struct{})

	errorsCount := 0

	wg.Add(len(tasks))
	for i := 0; i < goroutinesCount; i++ {
		go worker(tasksChan, errorsChan, successChan, wg)
	}

	for offset := 0; offset < len(tasks); offset += goroutinesCount {
		for i := 0; i < goroutinesCount && offset+i < len(tasks); i++ {
			tasksChan <- tasks[offset+i]
		}

		for i := 0; i < goroutinesCount && offset+i < len(tasks); i++ {
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

func worker(tasksChan <-chan task, errorsChan chan<- error, successChan chan<- struct{}, wg *sync.WaitGroup) {
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
