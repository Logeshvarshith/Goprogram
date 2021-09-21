package main

import (
	"fmt"
	"math"
)

func main() {
	isprime(29)

}

func isprime(n int) {
	if n <= 1 {
		fmt.Println("Number should be grater than 1")
		return
	}
	a := int(math.Sqrt(float64(n)))
	for i := 2; i <= a; i++ {
		if n%i == 0 {
			fmt.Println("Non prime number")
			return
		}
	}
	fmt.Print("prime number")
	return

}
