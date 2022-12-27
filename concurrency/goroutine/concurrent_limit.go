package goroutine

import (
	"fmt"
	"sync"
)

const defaultConcurrentLimit = 2

type ConcurrentLimiter struct {
	wg     sync.WaitGroup
	tokens chan struct{}
}

func NewConcurrentLimiter(limit int32) *ConcurrentLimiter {
	lmt := limit
	if limit <= 0 {
		lmt = defaultConcurrentLimit
	}
	return &ConcurrentLimiter{
		tokens: make(chan struct{}, lmt),
	}
}

func (cl *ConcurrentLimiter) setToken() {
	cl.tokens <- struct{}{}
}

func (cl *ConcurrentLimiter) backToken() {
	<-cl.tokens
}

func (cl *ConcurrentLimiter) incrCount() {
	cl.wg.Add(1)
}

func (cl *ConcurrentLimiter) decrCount() {
	cl.wg.Done()
}

func (cl *ConcurrentLimiter) Do(fun func() error) {
	cl.setToken()
	cl.incrCount()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover from Do, err: %v\n", r)
			}
		}()
		defer cl.backToken()
		defer cl.decrCount()

		fun()
	}()
}

func (cl *ConcurrentLimiter) Wait() {
	cl.wg.Wait()
}
