package tutorial02

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type greeting struct{}

type MyInt int

func (m MyInt) Greater(v MyInt) bool {
	return m > v
}

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

func TestIsPositive_IsTrueFor1(t *testing.T) {
	t.Parallel()

	input := MyInt(1)
	want := true
	got := IsPositive(input)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForNegative1(t *testing.T) {
	t.Parallel()

	input := MyInt(-1)
	want := false
	got := IsPositive(input)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositive_IsFalseForZero(t *testing.T) {
	t.Parallel()

	input := MyInt(0)
	want := false
	got := IsPositive(input)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsGreater_IsTrueFor2And1(t *testing.T) {
	t.Parallel()

	want := true
	got := IsGreater(MyInt(2), MyInt(1))

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsGreater_IsFalseFor1And2(t *testing.T) {
	t.Parallel()

	want := false
	got := IsGreater(MyInt(1), MyInt(2))

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
