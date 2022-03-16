package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	want := 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	want := 9
	got := calculator.Multiply(3, 3)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
