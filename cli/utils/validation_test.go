package utils

import (
	"fmt"
	"testing"
)

func TestValidator(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"testproject", true},
		{"test-project", true},
		{"test_project", true},
		{"test project", false},
		{" testproject", false},
		{"testproject ", false},
		{"te!stproject", false},
		{"testpro/ject", false},
		{"testproject%", false},
		{"*testproject", false},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("%s should be %t", testCase.input, testCase.expected)
		t.Run(testName, func(t *testing.T) {
			output := IsValidIdentifierString(testCase.input)

			if output != testCase.expected {
				t.Errorf("For %s got %t, expected %t", testCase.input, output, testCase.expected)
			}
		})
	}
}
