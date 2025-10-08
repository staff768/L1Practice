package main

import (
	"fmt"
	"math"
)

func bucket10(t float64) int {
	return int(math.Trunc(t/10.0)) * 10
}

func main() {
	values := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, v := range values {
		k := bucket10(v)
		groups[k] = append(groups[k], v)
	}

	for k, vs := range groups {
		fmt.Printf("%d: %v\n", k, vs)
	}
}
