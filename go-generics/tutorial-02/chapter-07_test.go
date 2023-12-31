package tutorial02

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPushOneValueToEmptyStackLeavesTheStackWithLength1(t *testing.T) {
	t.Parallel()
	s := Stack[int]{}
	s.Push(0)
	if s.Len() != 1 {
		t.Fatal("Push didn't add value to stack")
	}
}

func TestPushTwiceThenPopTwiceOnEmptyStackLeavesTheStackEmpty(t *testing.T) {
	t.Parallel()
	s := Stack[string]{}
	s.Push("a", "b")
	if s.Len() != 2 {
		t.Fatal("Push didn't add all values to stack")
	}
	s.Pop()
	want := "a"
	got, ok := s.Pop()
	if !ok {
		t.Fatal("Pop returned not ok on non-empty stack")
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestPopReturnsNotOKOnEmptyStack(t *testing.T) {
	t.Parallel()
	s := Stack[int]{}
	_, ok := s.Pop()
	if ok {
		t.Fatal("Pop returned ok on empty stack")
	}
}
