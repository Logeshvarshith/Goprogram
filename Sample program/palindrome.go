package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(ispalindrome("asist"))
}

func ispalindrome(s string) bool {

	upper := strings.ToUpper(s)

	for i := 0; i < len(upper)/2; i++ {
		if upper[i] != upper[len(upper)-i-1] {
			return false
		}

	}
	return true

}
