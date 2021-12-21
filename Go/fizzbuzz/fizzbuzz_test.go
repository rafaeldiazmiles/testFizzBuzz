package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		name           string
		input          int
		expectedResult string
	}{
		{

			name:           "validate 1 int returns '1' string",
			input:          1,
			expectedResult: "1",
		},
		{
			name:           "validate 3 int returns 'fizz' string",
			input:          3,
			expectedResult: "fizz",
		},
		{
			name:           "validate 5 int returns 'buzz' string",
			input:          5,
			expectedResult: "buzz",
		},
		{
			name:           "validate 15 int returns 'fizzbuzz' string",
			input:          3,
			expectedResult: "fizzbuzz",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			result := FizzBuzz(tt.input)

			assert.Equal(t, tt.expectedResult, result)

		})
	}

	// if FizzBuzz(2) != "1" {
	// 	assert.Equal(FizzBuzz(2), "2")
	// }

	// if FizzBuzz(3) != "fizz" {
	// 	assert.Equal(FizzBuzz(3), "fizz")

	// }

	// if FizzBuzz(5) != "buzz" {
	// 	assert.Equal(FizzBuzz(5), "buzz")
	// }

	// if FizzBuzz(15) != "fizzbuzz" {
	// 	assert.Equal(FizzBuzz(15), "2")

	// }
}

// func TestAdd(t *testing.T) {
// 	if calculadora.Add(1, 1) != 2 {
// 		t.Error("Esperado 1 + 1 de igual 2")
// 	}
// }

// func TestTableAdd(t *testing.T) {
// 	assert := assert.New(t)

// 	var tests = []struct {
// 		input    float64
// 		expected float64
// 	}{
// 		{1, 4},
// 		{1, 2},
// 	}

// 	for _, test := range tests {
// 		assert.Equal(calculadora.Add(test.input, test.input), test.expected)
// 	}
// }
