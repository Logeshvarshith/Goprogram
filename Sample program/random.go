package main

import (
	"fmt"
	"math/rand"
	"time"
)

var randowCodes = [...]byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
}

func main() {
	var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i <= 0; i++ {
		var pwd []byte = make([]byte, 6)

		for j := 0; j < 6; j++ {
			index := r.Int() % len(randowCodes)

			pwd[j] = randowCodes[index]
		}

		fmt.Printf("%s\n", string(pwd))
	}
}
