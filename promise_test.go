package abuse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fooPromise() *Promise[string] {
	return New(func() string {
		return "foo"
	})
}

func barPromise() *Promise[string] {
	return New(func() string {
		return "bar"
	})
}

func TestSingle(t *testing.T) {
	assert.Equal(t, "foo", fooPromise().Result())
}

func TestThen(t *testing.T) {
	result := fooPromise().
		Then(func(arg string) string {
			return arg + "bar"
		}).Result()

	assert.Equal(t, "foobar", result)
}

func TestPanic(t *testing.T) {
	p := New(func() string {
		panic("aaa")
	})

	result, exception := p.Await()

	assert.Equal(t, "", result)
	assert.Equal(t, "aaa", exception)
}

func TestElse(t *testing.T) {
	action1 := false
	action2 := false
	action3 := false
	action4 := false
	elseFlag := false

	chain := New(func() string {
		action1 = true
		return "action1"
	}).Then(func(string) string {
		action2 = true
		return "action2"
	}).Then(func(string) string {
		action3 = true
		panic("aaa")
	}).Then(func(string) string {
		action4 = true
		return "action4"
	}).Else(func(interface{}) string {
		elseFlag = true
		return "else"
	})

	result, exception := chain.Await()

	assert.Equal(t, "else", result, "result")
	assert.Equal(t, nil, exception, "exception")

	assert.Equal(t, true, action1, "action1")
	assert.Equal(t, true, action2, "action2")
	assert.Equal(t, true, action3, "action3")
	assert.Equal(t, false, action4, "action4")
	assert.Equal(t, true, elseFlag, "elseFlag")
}
