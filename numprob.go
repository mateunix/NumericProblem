package main

import (
	"fmt"
)

// Make a program that show Ping for all numbers between 1 and 100 that are divisible by 3.
// Pong for all numbers between 1 and 100 that are divisible by 5.
// PingPong  for all numbers between 1 and 100 that are divisible by 3 and 5 at the same time.

func main() {
	var x []string
	for i := 0; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			x = append(x, "PingPong")

		} else if i%3 == 0 && i%5 == 0 {
			x = append(x, "Ping")

		} else if i%5 == 0 {
			x = append(x, "Pong")

		}

	}

	fmt.Println(x)
}
