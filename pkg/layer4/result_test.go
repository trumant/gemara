package layer4

import (
	"testing"
)

func TestResultString(t *testing.T) {
	tests := []struct {
		name     string
		result   Result
		expected string
	}{
		{
			result:   Passed,
			expected: "Passed",
		},
		{
			result:   Failed,
			expected: "Failed",
		},
		{
			result:   NeedsReview,
			expected: "Needs Review",
		},
		{
			result:   Unknown,
			expected: "Unknown",
		},
		{
			result:   Result(4),
			expected: "Unknown",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.result.String()
			if actual != test.expected {
				t.Errorf("expected %q, got %q", test.expected, actual)
			}
		})
	}
}
