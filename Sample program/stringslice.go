package main

import "fmt"

func main() {
	a := []string{"abc", "cvd", "vdf", "ngs"}
	fmt.Print(a[0:3])
	fmt.Print(a[1:4])
	fmt.Print(a[0:])
	fmt.Print(a[:4])
}
