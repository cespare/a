package a_test

import (
	"testing"

	"github.com/cespare/a"
)

func TestMessage(t *testing.T) {
	ok, msg := a.Equals(1, 2, "mymessage")
	a.Check(t, ok, a.IsFalse)
	a.Check(t, "mymessage", a.Equals, msg)
}

func TestEquals(t *testing.T) {
	ok, _ := a.Equals(1, 1)
	a.Check(t, ok, a.IsTrue)
	ok, _ = a.Equals(1, 2)
	a.Check(t, ok, a.IsFalse)
	ok, _ = a.Equals(1, "two")
	a.Check(t, ok, a.IsFalse)
	ok, _ = a.Equals([]int{1, 2, 3}, []int{1, 2, 3})
	a.Check(t, ok, a.IsFalse)
}

func TestDeepEquals(t *testing.T) {
	ok, _ := a.DeepEquals(1, 1)
	a.Check(t, ok, a.IsTrue)
	ok, _ = a.DeepEquals(1, 2)
	a.Check(t, ok, a.IsFalse)
	ok, _ = a.DeepEquals(1, "two")
	a.Check(t, ok, a.IsFalse)
	ok, _ = a.DeepEquals([]int{1, 2, 3}, []int{1, 2, 3})
	a.Check(t, ok, a.IsTrue)
}

func TestIsNil(t *testing.T) {
	a.Check(t, nil, a.IsNil)
	var n *struct{}
	a.Check(t, n, a.IsNil)
	n = new(struct{})
	a.Check(t, n, a.NotNil)
}

func TestBool(t *testing.T) {
	a.Check(t, true, a.IsTrue)
	a.Check(t, false, a.IsFalse)
}
