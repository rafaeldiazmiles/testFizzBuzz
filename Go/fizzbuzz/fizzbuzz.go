package fizzbuzz

import "strconv"

func FizzBuzz(x int) string {

	if x%3 == 0 && x%5 == 0 {
		return "fizzbuzz"
	} else if x%3 == 0 {
		return "fizz"
	} else if x%5 == 0 {
		return "buzz"
	} else {
		ret := strconv.Itoa(x)
		return ret
	}
}
