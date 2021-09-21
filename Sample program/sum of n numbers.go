package main

import "fmt"

var y int

func main() {
	y := 0
	for i := 1; i <= 10; i++ {
		y = i + y

	}
	fmt.Print(y)
}
