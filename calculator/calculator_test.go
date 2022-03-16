package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	want := 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
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
