package tutorial02

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestProductOfInts2And3Is6(t *testing.T) {
	t.Parallel()
	want := 6
	got := Product(2, 3)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestProductOfFloats1Point6And2Point3Is3Point68(t *testing.T) {
	t.Parallel()
	want := 3.68
	got := Product(1.6, 2.3)
	if !cmp.Equal(want, got, cmpopts.EquateApprox(0.0001, 0)) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestProductOfComplex2Plus3iAnd3Plus2iIs0Plus13i(t *testing.T) {
	t.Parallel()
	want := 0 + 13i
	got := Product(2+3i, 3+2i)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupesIsTrueWhenInputContainsNonConsecutiveDuplicates(t *testing.T) {
	t.Parallel()
	s := []int{1, 2, 3, 1, 5}
	want := true
	got := Dupes(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupesIsFalseWhenInputContainsNoDuplicates(t *testing.T) {
	t.Parallel()
	s := []string{"a"}
	want := false
	got := Dupes(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
