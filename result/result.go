package result

import (
	"fmt"
	. "github.com/initdc/type/option"
)

type Result[T any, E any] struct {
	value    T
	err      E
	ok       bool
	assigned bool
}

func Ok[T any, E any](v T) Result[T, E] {
	return Result[T, E]{
		value:    v,
		ok:       true,
		assigned: true,
	}
}

func Err[T any, E any](e E) Result[T, E] {
	return Result[T, E]{
		err:      e,
		ok:       false,
		assigned: true,
	}
}

func (r *Result[T, E]) Ok(v T) {
	if r.assigned {
		panic("Result already assigned")
	}
	r.value = v
	r.ok = true
	r.assigned = true
}

func (r *Result[T, E]) Err(e E) {
	if r.assigned {
		panic("Result already assigned")
	}
	r.err = e
	r.ok = false
	r.assigned = true
}

func (r *Result[T, E]) IsOk() bool {
	return r.ok
}

func (r *Result[T, E]) IsOkAnd(f func(T) bool) bool {
	if !r.ok {
		return false
	}
	return f(r.value)
}

func (r *Result[T, E]) IsErr() bool {
	return !r.ok
}

func (r *Result[T, E]) IsErrAnd(f func(E) bool) bool {
	if r.ok {
		return false
	}
	return f(r.err)
}

func (r *Result[T, E]) ConvOk() Option[T] {
	if r.ok {
		return Some(r.value)
	}
	return None[T]()
}

func (r *Result[T, E]) ConvErr() Option[E] {
	if r.ok {
		return None[E]()
	}
	return Some(r.err)
}

func (r *Result[T, E]) Expect(msg string) T {
	if r.ok {
		return r.value
	}
	panic(fmt.Sprintf("%s: %v", msg, r.err))
}

func (r Result[T, E]) Unwrap() T {
	if r.ok {
		return r.value
	}
	panic(fmt.Sprintf("called Result.Unwrap() on an Err value: %v", r.err))
}

func (r *Result[T, E]) UnwrapOrDefault() T {
	if r.ok {
		return r.value
	}
	var zero T
	return zero
}

func (r *Result[T, E]) ExpectErr(msg string) E {
	if r.ok {
		panic(fmt.Sprintf("%s: %v", msg, r.value))
	}
	return r.err
}

func (r *Result[T, E]) UnwrapErr() E {
	if r.ok {
		panic(fmt.Sprintf("called Result.UnwrapErr() on an Ok value: %v", r.value))
	}
	return r.err
}

func (r *Result[T, E]) UnwrapOr(def T) T {
	if r.ok {
		return r.value
	}
	return def
}

func (r *Result[T, E]) UnwrapOrElse(f func(E) T) T {
	if r.ok {
		return r.value
	}
	return f(r.err)
}
