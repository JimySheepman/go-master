package tutorial02

import "testing"

func TestEmptyIsTrueForEmptySequence(t *testing.T) {
	t.Parallel()

	s := Sequence[int]{}
	got := s.Empty()

	if !got {
		t.Fatal("false for empty sequence")
	}
}

func TestEmptyIsFalseForNonEmptySequence(t *testing.T) {
	t.Parallel()

	s := Sequence[string]{"a", "b", "c"}
	got := s.Empty()

	if got {
		t.Fatal("false for empty sequence")
	}
}
