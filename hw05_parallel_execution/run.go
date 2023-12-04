package hw05parallelexecution

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type ErrCounter struct {
	mx sync.Mutex
	m  int32
}

func (c *ErrCounter) Reset() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m = 0
}

func (c *ErrCounter) Inc() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m++
}

// Compare - возвращает true, если количество ошибок больше заданного числа, иначе false
// Всегда возвращает false если n<0,т.е. игнорирование ошибок.
func (c *ErrCounter) Compare(n int32) bool {
	if n < 0 {
		return false
	}
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.m >= n
}

var errCount ErrCounter

// Run starts tasks in n goroutines
// and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	errCount.Reset()
	job := make(chan Task, n/2)
	ctx, ctxCancel := context.WithCancel(context.Background())

	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(ctx, &wg, job)
	}

	wg.Add(1)
	go writeTask(ctx, &wg, job, tasks)

	go func(ctx context.Context, ctxCancel context.CancelFunc) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if errCount.Compare(int32(m)) {
					ctxCancel()
					return
				}
			}
		}
	}(ctx, ctxCancel)

	wg.Wait()
	ctxCancel()

	if errCount.Compare(int32(m)) {
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
		default:
			job <- task
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, job <-chan Task) {
	defer wg.Done()
	for task := range job {
		select {
		case <-ctx.Done():
			return
		default:
			err := task()
			if err != nil {
				errCount.Inc()
			}
		}
		time.Sleep(1 * time.Millisecond)
	}
}
