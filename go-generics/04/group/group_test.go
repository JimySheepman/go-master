package group_test

import (
	"go-generics/04/group"
	"slices"
	"testing"
)

func TestGroupContainsWhatIsAppendedToIt(t *testing.T) {
	t.Parallel()
	got := group.Group[string]{}
	got = append(got, "hello")
	got = append(got, "world")
	want := group.Group[string]{"hello", "world"}
	if !slices.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
