package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		number int
		expected string
	}{
		{0, "Even"},
		{1, "Odd"},
		{2, "Even"},
		{3, "Odd"},
	}
	for _, test := range tests {
		result := EvenOrOdd(test.number)
		if result != test.expected {
			t.Errorf("EvenOrOdd(%d) = %s; want %s", test.number, result, test.expected)
		}
	}
}