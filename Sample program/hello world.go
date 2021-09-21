package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// add 45 minutes
	timeout := now.Add(time.Minute * 1)
	fmt.Print(timeout)
}
