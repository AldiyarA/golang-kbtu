package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, ErrNegativeSqrt(x)
	}
	if x == 0 {
		return 0, nil
	}
	z := float64(1)
	y := z
	z -= (z*z - x) / (2 * z)
	i := 0
	for math.Abs(x-y*y) > math.Abs(x-z*z) || i <= 5 {
		i++
		y = z
		z -= (z*z - x) / (2 * z)
	}
	return y, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(-2))
}
