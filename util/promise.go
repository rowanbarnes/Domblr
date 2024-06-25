package util

import (
	"sync"
)

type Promise struct {
	resultChan chan []any
	errorChan  chan error
	wg         sync.WaitGroup
}

func NewPromise() *Promise {
	return &Promise{
		resultChan: make(chan []any),
		errorChan:  make(chan error),
	}
}

func (p *Promise) Then(callback func([]any)) *Promise {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		result := <-p.resultChan
		callback(result)
	}()
	return p
}

func (p *Promise) Catch(callback func(error)) *Promise {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		err := <-p.errorChan
		callback(err)
	}()
	return p
}

func (p *Promise) Resolve(result []any) {
	println("Promise resolved")
	p.resultChan <- result
}

func (p *Promise) Reject(err error) {
	if err != nil {
		println("Promise rejected: " + err.Error())
		p.errorChan <- err
	}
}
