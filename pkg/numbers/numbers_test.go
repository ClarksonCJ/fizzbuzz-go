package numbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByThree(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 1,
			want:  false,
		},
		{
			input: 2,
			want:  false,
		},
		{
			input: 3,
			want:  true,
		},
		{
			input: 30,
			want:  true,
		},
	}
	processor := &Processor{}
	for _, tt := range tests {
		result := processor.ByThree(tt.input)
		assert.Equal(tt.want, result, "Expected result to match Wanted")
	}
}

func TestByFive(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 1,
			want:  false,
		},
		{
			input: 2,
			want:  false,
		},
		{
			input: 3,
			want:  false,
		},
		{
			input: 30,
			want:  true,
		},
		{
			input: 100,
			want:  true,
		},
	}
	processor := &Processor{}
	for _, tt := range tests {
		result := processor.ByFive(tt.input)
		assert.Equal(tt.want, result, "Expected result to match Wanted")
	}
}

func TestByThreeAndFive(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 1,
			want:  false,
		},
		{
			input: 3,
			want:  false,
		},
		{
			input: 30,
			want:  true,
		},
		{
			input: 15,
			want:  true,
		},
		{
			input: 100,
			want:  false,
		},
	}
	processor := &Processor{}
	for _, tt := range tests {
		result := processor.ByThreeAndFive(tt.input)
		assert.Equal(tt.want, result, fmt.Sprintf("Expected result of %d to match Wanted", tt.input))
	}
}
