package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines
// and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var errCount int32
	atomic.StoreInt32(&errCount, 0)

	job := make(chan Task, n)
	ctx, ctxCancel := context.WithCancel(context.Background())

	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(ctxCancel, &wg, job, &errCount, int32(m))
	}

	wg.Add(1)
	go writeTask(ctx, &wg, job, tasks)

	wg.Wait()
	ctxCancel()

	if m >= 0 && atomic.LoadInt32(&errCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func writeTask(ctx context.Context, wg *sync.WaitGroup, job chan<- Task, tasks []Task) {
	defer wg.Done()
	defer close(job)
	for _, task := range tasks {
		select {
		case <-ctx.Done():
			return
		case job <- task:
		}
	}
}

func worker(ctxCancel context.CancelFunc, wg *sync.WaitGroup, job <-chan Task, errCounter *int32, maxError int32) {
	defer wg.Done()
	for task := range job {
		err := task()
		if err != nil && maxError >= 0 {
			atomic.AddInt32(errCounter, 1)
			if atomic.LoadInt32(errCounter) >= maxError {
				ctxCancel()
				return
			}
		}
	}
}
