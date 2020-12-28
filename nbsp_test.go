package nbsp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -run TestCS
func TestCS(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		description string
	}{
		{
			"",
			"",
			"empty string",
		},
		{
			" ",
			" ",
			"one space",
		},
		{
			"  ",
			"  ",
			"two spaces",
		},
		{
			"   ",
			"   ",
			"three spaces",
		},
		{
			"1 ",
			"1\u00A0",
			"one digit with a space",
		},
		{
			"12 ",
			"12\u00A0",
			"two digits with a space",
		},
		{
			"12. ",
			"12.\u00A0",
			"two digits with a dot and space",
		},
		{
			"Ahoj",
			"Ahoj",
			"word starting with an uppercase",
		},
		{
			"Ahoj ",
			"Ahoj ",
			"word starting with an uppercase, followed by a space",
		},
		{
			"Ahoj Franto",
			"Ahoj Franto",
			"two words with a space between",
		},
		{
			"A potom taky.",
			"A\u00A0potom taky.",
			"simple sentence with starting with an uppercase",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, Replace(test.input, CS), test.description)
	}
}
