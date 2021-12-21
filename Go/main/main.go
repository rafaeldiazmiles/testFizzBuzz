package main

import (
	"fizzbuzz"
	"fmt"
)

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("Input: %v Output: %v\n", i, fizzbuzz.FizzBuzz(i))
	}
}
