package q12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	input  int
	output string
}{
	{
		input:  3,
		output: "III",
	},
	{
		input:  4,
		output: "IV",
	},
	{
		input:  9,
		output: "IX",
	},
	{
		input:  58,
		output: "LVIII",
	},
	{
		input:  1994,
		output: "MCMXCIV",
	},
	{
		input:  3999,
		output: "MMMCMXCIX",
	},
	{
		input:  6,
		output: "VI",
	},
}

func TestIntToRomanMap(t *testing.T) {
	for i := range cases {
		actual := intToRomanMap(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}

func TestIntToRoman(t *testing.T) {
	for i := range cases {
		actual := intToRoman(cases[i].input)
		assert.Equal(t, cases[i].output, actual)
	}
}
