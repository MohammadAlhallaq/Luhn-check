# LuhnAlgo

This is a Go package that provides functionality for validating credit card numbers using the [Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm).

## Installation

To use this package, you can simply import it in your Go project:

```go
import "github.com/MohammadAlhallaq/luhnAlgo"

func main() {
	luhn.Valid(24354234) //= false
	luhn.Valid(40000070302000001) //= true
}