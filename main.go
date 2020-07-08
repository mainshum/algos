package main

import (
	"fmt"
	"math/big"

	"github.com/mainshum/algos/A1_W1"
)

func main() {
	num1 := "3141592653589793238462643383279502884197169399375105820974944592"
	num2 := "2718281828459045235360287471352662497757247093699959574966967627"

	var x, _ = new(big.Int).SetString(num1, 10)
	var y, _ = new(big.Int).SetString(num2, 10)

	result := A1_W1.Karatsuba(*x, *y)

	fmt.Println(result.String())
}
