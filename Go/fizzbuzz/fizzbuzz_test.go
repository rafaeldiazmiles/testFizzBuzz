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
			input:          15,
			expectedResult: "fizzbuzz",
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {

			result := FizzBuzz(tc.input)

			assert.Equal(tc.expectedResult, result)

		})
	}
}
