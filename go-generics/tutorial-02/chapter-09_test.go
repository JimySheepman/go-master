package tutorial02

import (
	"github.com/google/go-cmp/cmp"
	"testing"
	"unicode"
)

func positive(p int) bool {
	return p > 0
}

func TestContainsFunc_IsTrueForPositiveOnInputContainingPositiveInts(t *testing.T) {
	t.Parallel()
	input := []int{-2, 0, 1, -1, 5}
	got := ContainsFunc(input, positive)
	if !got {
		t.Fatalf("%v: want true for 'contains positive', got false", input)
	}
}

func TestContainsFunc_IsFalseForIsUpperOnInputContainingNoUppercaseRunes(t *testing.T) {
	t.Parallel()
	input := []rune("hello")
	got := ContainsFunc(input, unicode.IsUpper)
	if got {
		t.Fatalf("%q: want false for 'contains uppercase', got true", input)
	}
}

func TestMergeCorrectlyMergesTwoMapsOfIntToBool(t *testing.T) {
	t.Parallel()
	inputs := []map[int]bool{
		{
			1: false,
			2: false,
			3: false,
		},
		{
			3: true,
			5: true,
		},
	}
	want := map[int]bool{
		1: false,
		2: false,
		3: true,
		5: true,
	}
	got := Merge(inputs...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestMergeCorrectlyMergesThreeMapsOfStringToAny(t *testing.T) {
	t.Parallel()
	inputs := []map[string]any{
		{
			"a": nil,
		},
		{
			"b": "hello, world",
			"c": 0,
		},
		{
			"a": 6 + 2i,
		},
	}
	want := map[string]any{
		"a": 6 + 2i,
		"b": "hello, world",
		"c": 0,
	}
	got := Merge(inputs...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
