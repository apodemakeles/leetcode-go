package q20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	s      string
	output bool
}{
	//{
	//	s:      "()",
	//	output: true,
	//},
	//{
	//	s:      "()[]{}",
	//	output: true,
	//},
	//{
	//	s:      "()[]{}",
	//	output: true,
	//},
	//{
	//	s:      "(]",
	//	output: false,
	//},
	//{
	//	s:      "([)]",
	//	output: false,
	//},
	//{
	//	s:      "{[]}",
	//	output: true,
	//},
	{
		s:      "((",
		output: false,
	},
}

func TestIsValid(t *testing.T) {
	for i := range cases {
		actual := isValid(cases[i].s)
		assert.Equal(t, cases[i].output, actual)
	}
}
