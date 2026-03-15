package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string // A description of the test case
		a        int    // Input A
		b        int    // Input B
		expected int    // Expected result
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -10, 5, -5},
		{"zeros", 0, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
