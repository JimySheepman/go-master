package tutorial02

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
	"unicode"
)

func TestFuncMap_AppliesDoubleTo2AndReturns4(t *testing.T) {
	t.Parallel()

	fm := FuncMap[int, int]{
		"double": func(i int) int {
			return i * 2
		},
		"addOne": func(i int) int {
			return i + 1
		},
	}

	want := 4
	got := fm.Apply("double", 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFuncMap_AppliesUpperToUppercaseInputAndReturnsTrue(t *testing.T) {
	t.Parallel()

	fm := FuncMap[rune, bool]{
		"upper": unicode.IsUpper,
		"lower": unicode.IsLower,
	}

	got := fm.Apply("upper", 'A')
	if !got {
		t.Fatal("upper('A'): want true, got false")
	}
}

func isOdd(p int) bool {
	return p%2 != 0
}

func next(p int) int {
	return p + 1
}

func TestComposeAppliesFuncsToIntInReverseOrder(t *testing.T) {
	t.Parallel()
	odd := Compose(isOdd, next, 1)
	if odd {
		t.Fatal("isOdd(next(1)): want false, got true")
	}
}

func TestComposeAppliesFuncsToStringInReverseOrder(t *testing.T) {
	t.Parallel()
	input := "HeLlO, wOrLd"
	want := "hello, world"
	got := Compose(strings.ToLower, strings.ToUpper, input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func first[E any](s []E) E {
	return s[0]
}

func last[E any](s []E) E {
	return s[len(s)-1]
}

func TestComposeAppliesFuncsToSliceInReverseOrder(t *testing.T) {
	t.Parallel()
	input := [][]int{{1, 2, 3}}
	want := 1
	got := Compose(first[int], last[[]int], input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestComposeAppliesFuncsToSliceInReverseOrder2(t *testing.T) {
	t.Parallel()
	input := [][]int{{1, 2, 3}}
	want := 3
	got := Compose(last[int], first[[]int], input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
