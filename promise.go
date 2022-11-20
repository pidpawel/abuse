package abuse

import (
	"sync"
)

type Promise[T any] struct {
	result    T
	exception interface{}

	callable func() T
	isLazy   bool

	once sync.Once
	lock sync.Mutex
	wait sync.WaitGroup
}

func (p *Promise[T]) Do() {
	p.once.Do(func() {
		go func() {
			defer func() {
				if exception := recover(); exception != nil {
					p.lock.Lock()
					defer p.lock.Unlock()

					p.exception = exception

					p.wait.Done()
				}
			}()

			result := p.callable()

			p.lock.Lock()
			defer p.lock.Unlock()

			p.result = result

			p.wait.Done()
		}()
	})
}

func newInternal[T any](callable func() T, isLazy bool) *Promise[T] {
	p := Promise[T]{
		callable: callable,
		isLazy:   isLazy,
	}

	p.wait.Add(1)

	if !p.isLazy {
		p.Do()
	}

	return &p
}

func New[T any](callable func() T) *Promise[T] {
	return newInternal(callable, true)
}

func NewLazy[T any](callable func() T) *Promise[T] {
	return newInternal(callable, false)
}

func NewChild[T any](parent *Promise[T], callable func() T) *Promise[T] {
	var resultingPromise *Promise[T]

	// If the original one is lazy, than the whole chain should be lazy, right?
	if parent.IsLazy() {
		resultingPromise = NewLazy(callable)
	} else {
		resultingPromise = New(callable)
	}

	return resultingPromise
}

func (p *Promise[T]) IsLazy() bool {
	return p.isLazy
}

func (p *Promise[T]) Await() (T, interface{}) {
	p.Do()
	p.wait.Wait()

	return p.result, p.exception
}

func (p *Promise[T]) Result() T {
	result, exception := p.Await()
	if exception != nil {
		panic(exception)
	}

	return result
}

func (p *Promise[T]) Exception() interface{} {
	_, exception := p.Await()
	return exception
}

func (p *Promise[T]) Then(callback func(T) T) *Promise[T] {
	return NewChild(p, func() T {
		return callback(p.Result())
	})
}

func (p *Promise[T]) Else(callback func(interface{}) T) *Promise[T] {
	return NewChild(p, func() T {
		result, exception := p.Await()
		if exception != nil {
			return callback(exception)
		}

		return result
	})
}
