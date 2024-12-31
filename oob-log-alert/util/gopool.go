package util

import (
	"sync"
)

// Pool for control goroutine num to run
type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// NewPool make a pool with size
func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add a goroutine to pool with delta num
func (p *Pool) Add() {
	p.queue <- 1
	p.wg.Add(1)
}

// Done invoke this while goroutine is done
func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

// Wait wait all goroutine is finished
func (p *Pool) Wait() {
	p.wg.Wait()
}
