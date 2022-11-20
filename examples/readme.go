package main

import (
	"fmt"

	"github.com/pidpawel/abuse"
)

func main() {
	result, exception := (&abuse.PromiseSlice[int]{
		abuse.New(func() int {
			return 1
		}),
		abuse.New(func() int {
			return 2
		}),
		abuse.New(func() int {
			return 3
		}),
	}).ForEach(func(arg int) int {
		return 10 + arg
	}).Fold(0, func(a, b int) int {
		return a + b
	}).Await()

	if exception != nil {
		panic(exception)
	}

	fmt.Println(result) // Should be 11 + 12 + 13 = 36
}
