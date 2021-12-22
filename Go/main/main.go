package main

import (
	"fmt"

	"github.com/rafaeldiazmiles/testFizzBuzz/fizzbuzz"
)

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("Input: %v Output: %v\n", i, fizzbuzz.FizzBuzz(i))
	}
}
