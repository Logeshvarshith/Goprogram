package main

import "fmt"

func main() {
	d := []string{"hello", "hai", "Love"}

	for i, c := range d {
		fmt.Print(i, c)
	}
}
