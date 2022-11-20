package abuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceSimple(t *testing.T) {
	slice := PromiseSlice[string]{
		fooPromise(),
		barPromise(),
	}

	results, exceptions := slice.Await()

	assert.Equal(t, []string{"foo", "bar"}, results)
	assert.Equal(t, []interface{}{nil, nil}, exceptions)
}

func TestForEach(t *testing.T) {
	slice := (&PromiseSlice[string]{
		fooPromise(),
		barPromise(),
	}).ForEach(func(arg string) string {
		return arg + "!"
	})

	results, exceptions := slice.Await()

	assert.Equal(t, []string{"foo!", "bar!"}, results)
	assert.Equal(t, []interface{}{nil, nil}, exceptions)
}

func TestFold(t *testing.T) {
	result, exception := (&PromiseSlice[string]{
		fooPromise(),
		barPromise(),
	}).Fold("", func(a, b string) string {
		return a + " " + b
	}).Await()

	assert.Equal(t, " foo bar", result)
	assert.Equal(t, nil, exception)
}
