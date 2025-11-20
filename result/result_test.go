package result

import (
	"testing"
	. "github.com/initdc/type/option"
)

func TestIsOk(t *testing.T) {
	r := Ok[int, string](1)
	if r.IsOk() == false {
		t.Error("TestIsOk case 1 failed")
	}

	e := Err[int, string]("error")
	if e.IsOk() == true {
		t.Error("TestIsOk case 2 failed")
	}
}

func TestIsErr(t *testing.T) {
	var r Result[int, string]
	r.Ok(1)

	if r.IsErr() == true {
		t.Error("TestIsErr case 1 failed")
	}

	var e Result[int, string]
	e.Err("error")
	if e.IsErr() == false {
		t.Error("TestIsErr case 2 failed")
	}
}


func TestIsOkAnd(t *testing.T) {
	r := Ok[int, string](1)
	if r.IsOkAnd(func(i int) bool { return i == 1}) == false {
		t.Error("TestOkAnd case 1 failed")
	}

	e := Err[int, string]("error")
	if e.IsOkAnd(func(i int) bool { return i == 1}) == true {
		t.Error("TestOkAnd case 2 failed")
	}
}

func TestIsErrAnd(t *testing.T) {
	r := Ok[int, string](1)
	if r.IsErrAnd(func(i string) bool { return i == "error"}) == true {
		t.Error("TestIsErrAnd case 1 failed")
	}

	e := Err[int, string]("error")
	if e.IsErrAnd(func(i string) bool { return i == "error"}) == false {
		t.Error("TestIsErrAnd case 2 failed")
	}
}

func TestConvOk(t *testing.T) {
	r := Ok[int, string](1)
	if r.ConvOk() != Some(1) {
	  t.Error("TestConvOk case 1 failed")
	}

	e := Err[int, string]("error")
	if e.ConvOk() != None[int]() {
		t.Error("TestConvOk case 2 failed")
	}
}

func TestConvErr(t *testing.T) {
	r := Ok[int, string](1)
	if r.ConvErr() != None[string]() {
		t.Error("TestConvErr case 1 failed")
	}

	e := Err[int, string]("error")
	if e.ConvErr() != Some("error") {
		t.Error("TestConvErr case 2 failed")
	}
}

func TestUnwrap(t *testing.T) {
	r := Ok[int, string](1)
	if r.Unwrap() != 1 {
		t.Error("TestUnwrap case 1 failed")
	}

	e := Err[int, string]("error")
	if e.UnwrapErr() != "error" {
		t.Error("TestUnwrap case 2 failed")
	}
}


func TestUnwrapErr(t *testing.T) {
	e := Err[int, string]("error")
	if e.UnwrapErr() != "error" {
		t.Error("TestUnwrapErr case 1 failed")
	}
}


func TestUnwrapOr(t *testing.T) {
	r := Ok[int, string](1)
	if r.UnwrapOr(2) != 1 {
		t.Error("TestUnwrapOr case 1 failed")
	}

	e := Err[int, string]("error")
	if e.UnwrapOr(2) != 2 {
		t.Error("TestUnwrapOr case 2 failed")
	}
}

func TestUnwrapOrElse(t *testing.T) {
	r := Ok[int, string](1)
	if r.UnwrapOrElse(func(e string) int { return 2 }) != 1 {
		t.Error("TestUnwrapOrElse case 1 failed")
	}

	e := Err[int, string]("error")
	if e.UnwrapOrElse(func(e string) int { return 2 }) != 2 {
		t.Error("TestUnwrapOrElse case 2 failed")
	}
}

