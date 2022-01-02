# slice

slice is a simple Go package to provide generic versions of `Map`, `Reduce` and 
`Filter` on slices. I mainly wrote it as an exercise to get more familar with 
Go 1.18's new type parameterization feature, but decided to publish it in case 
it is useful for others.

```go
package main

import (
	"fmt"
	"math"

	"github.com/akrennmair/slice"
)

func main() {
	someNumbers := []float64{1.0, 2.0, 4.0, 23.5, 42.9}
	fmt.Printf("sqrt(%v) = %v\n", someNumbers, slice.Map(someNumbers, func(v float64) float64 {
		return math.Sqrt(v)
	}))

	someStrings := []string{"hello", "to", "the", "golang", "community"}
	fmt.Printf("len(%v) = %v\n", someStrings, slice.Map(someStrings, func(s string) int {
		return len(s)
	}))
	fmt.Printf("total length: %d\n", slice.Reduce(slice.Map(someStrings, func(s string) int {
		return len(s)
	}), func(acc, i int) int {
		return acc + i
	}))

	moreNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("even numbers: %v\n", slice.Filter(moreNumbers, func(i int) bool {
		return i%2 == 0
	}))
}
```

## channel

An additional package `github.com/akrennmair/slice/channel` is also provided that contains
`Map`, `Reduce` and `Filter` functions that operate on `chan`s.

## Authors

* Andreas Krennmair <ak@synflood.at>

## License

See the file `LICENSE` for license information.
