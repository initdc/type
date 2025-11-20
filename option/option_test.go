package option

import (
	"testing"
)

func TestIsSome(t *testing.T) {
	s := Some(1)
	if s.IsSome() == false {
		t.Error("TestIsSome case 1 failed")
	}

	n := None[int]()
	if n.IsSome() == true {
		t.Error("TestIsSome case 2 failed")
	}
}

func TestIsNone(t *testing.T) {
	var s Option[int]
	s.Some(1)
	if s.IsNone() == true {
		t.Error("TestIsNone case 1 failed")
	}

	var n Option[int]
	n.None()
	if n.IsNone() == false {
		t.Error("TestIsNone case 2 failed")
	}
}

func TestIsSomeAnd(t *testing.T) {
	s := Some(1)
	if s.IsSomeAnd(func(int) bool { return true }) == false {
		t.Error("TestIsSomeAnd case 1 failed")
	}

	n := None[int]()
	if n.IsSomeAnd(func(int) bool { return true }) == true {
		t.Error("TestIsSomeAnd case 2 failed")
	}
}

func TestIsNoneOr(t *testing.T) {
	s := Some(1)
	if s.IsNoneOr(func(int) bool { return false }) == true {
		t.Error("TestIsNoneOr case 1 failed")
	}

	n := None[int]()
	if n.IsNoneOr(func(int) bool { return false }) == false {
		t.Error("TestIsNoneOr case 2 failed")
	}
}

func TestUnwrap(t *testing.T) {
	s := Some(1)
	if s.Unwrap() != 1 {
		t.Error("TestUnwrap case 1 failed")
	}

	n := None[int]()
	if n.UnwrapOr(99) != 99 {
		t.Error("TestUnwrap case 2 failed")
	}
}

func TestUnwrapOrElse(t *testing.T) {
	s := Some(1)
	if s.UnwrapOrElse(func() int { return 99 }) != 1 {
		t.Error("TestUnwrapOrElse case 1 failed")
	}

	n := None[int]()
	if n.UnwrapOrElse(func() int { return 99 }) != 99 {
		t.Error("TestUnwrapOrElse case 2 failed")
	}
}

func TestUnwrapOrDefault(t *testing.T) {
	s := Some(1)
	if s.UnwrapOrDefault() != 1 {
		t.Error("TestUnwrapOrDefault case 1 failed")
	}

	n := None[int]()
	if n.UnwrapOrDefault() != 0 {
		t.Error("TestUnwrapOrDefault case 2 failed")
	}
}
