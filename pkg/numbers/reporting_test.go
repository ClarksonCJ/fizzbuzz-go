package numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzReport(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input int
		want  string
	}{
		{
			input: 1,
			want:  "1",
		},
		{
			input: 2,
			want:  "2",
		},
		{
			input: 3,
			want:  "Fizz",
		},
		{
			input: 5,
			want:  "Buzz",
		},
		{
			input: 15,
			want:  "FizzBuzz",
		},
		{
			input: 30,
			want:  "FizzBuzz",
		},
		{
			input: 100,
			want:  "Buzz",
		},
	}
	reporter := &Reporter{}
	for _, tt := range tests {
		result := reporter.Report(tt.input)
		assert.Equal(tt.want, result, fmt.Sprintf("Expected Report for %d to be %s", tt.input, tt.want))
	}
}
