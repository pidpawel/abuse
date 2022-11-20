package abuse

type PromiseSlice[T any] []*Promise[T]

func (p *PromiseSlice[T]) Await() ([]T, []interface{}) {
	results := []T{}
	exceptions := []interface{}{}

	for _, promise := range *p {
		result, exception := promise.Await()

		results = append(results, result)
		exceptions = append(exceptions, exception)
	}

	return results, exceptions
}

func (p *PromiseSlice[T]) ForEach(callback func(T) T) *PromiseSlice[T] {
	promises := PromiseSlice[T]{}

	for _, promise := range *p {
		promises = append(promises, promise.Then(callback))
	}

	return &promises
}

func (p *PromiseSlice[T]) Fold(init T, callback func(T, T) T) *Promise[T] {
	return New(func() T {
		previous := init

		for _, promise := range *p {
			current := promise.Result()
			previous = callback(previous, current)
		}

		return previous
	})
}
