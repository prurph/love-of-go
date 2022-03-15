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
