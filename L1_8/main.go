package main

import (
	"errors"
	"fmt"
	"math/bits"
)

func setBit(x int64, i uint, value uint) (int64, error) {
	if value != 0 && value != 1 {
		return 0, errors.New("value must be 0 or 1")
	}
	if i >= uint(bits.UintSize*2) {
		return 0, errors.New("index out of range")
	}
	mask := int64(1) << i
	if value == 1 {
		return x | mask, nil
	}
	return x &^ mask, nil
}

func main() {

	res, err := setBit(5, 0, 0)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
