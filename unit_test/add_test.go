package unit_test

import (
	"testing"
)

var scenarios = []struct {
	name     string
	a, b     int
	expected int
}{
	{"positive numbers", 2, 3, 5},
	{"negative numbers", -2, -3, -5},
	{"mixed numbers", -2, 3, 1},
	{"zero", 0, 0, 0},
}

func TestAdd(t *testing.T) {
	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
