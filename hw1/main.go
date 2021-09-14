package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	y := z
	z -= (z*z - x) / (2 * z)
	i := 0
	for math.Abs(x-y*y) > math.Abs(x-z*z) || i <= 5 {
		i++
		y = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(y, z)
	}
	return y
}

func main() {
	fmt.Println(Sqrt(1001))
}
