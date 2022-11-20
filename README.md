# Abuse Golang library

## What?

```go
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
```

## How?

Generics. They're both a blessing, a curse and a mean to an end of this language. (This is not a place for such debate but tldr: they're powerful enough to enable people to port multiple concepts from other languages that may stand against core values of Go.)

Also goroutines and **no** reflection. That, unfortunately, doesn't mean that it will be blazing fast tough. Should be OK for IO/networking but awful for quick and dirty local patchups.

## Why?

Oh boy. I've stumbled upon [lo](https://github.com/samber/lo), [mo](https://github.com/samber/mo) and [promise](https://github.com/chebyrash/promise) while I was sick during time off and it seemed like a fitting way of spending an evening.

## Install

Don't.

Seriously though I'd be glad if someone used my work, but this particular piece seems to me more like a cheap way of achieveing a complicated result disregarding all the powerful tools that Go provides. If you find it fitting for your project despite this - sure, give it a go but remember that most likely one day you'll have to pay that technical debt.

## License

See `LICENSE.md`
