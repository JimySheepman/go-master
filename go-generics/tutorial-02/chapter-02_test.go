package tutorial02

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrintAnythingToPrintsInputToSuppliedWriter(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	PrintAnythingTo[string](buf, "Hello, world")
	want := "Hello, world\n"

	got := buf.String()

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGroupContainsWhatIsAppendedToIt(t *testing.T) {
	t.Parallel()

	got := Group[string]{}
	got = append(got, "hello")
	got = append(got, "world")

	want := Group[string]{"hello", "world"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestLenOfGroupIs2WhenItContains2Elements(t *testing.T) {
	t.Parallel()

	g := Group[int]{1, 2}
	want := 2

	got := Len(g)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
