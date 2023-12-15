package tutorial02

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type greeting struct{}

func (greeting) String() string {
	return "Howdy!"
}

func TestStringifyToPrintsResultOfStringMethodToSuppliedWriter(t *testing.T) {
	t.Parallel()

	buf := &bytes.Buffer{}
	StringifyTo[greeting](buf, greeting{})

	want := "Howdy!\n"

	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
