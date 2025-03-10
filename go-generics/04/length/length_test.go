package length_test

import (
	"go-generics/04/length"
	"testing"
)

func TestLenOfSliceIs2WhenItContains2Elements(t *testing.T) {
	t.Parallel()
	s := []int{1, 2}
	want := 2
	got := length.Len(s)
	if want != got {
		t.Errorf("Len(%v): want %d, got %d", s, want, got)
	}
}
