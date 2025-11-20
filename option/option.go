package option

type Option[T any] struct {
	value    T    `json:"value"`
	some     bool `json:"some"`
	assigned bool `json:"assigned"`
}

func Some[T any](v T) Option[T] {
	return Option[T]{
		value:    v,
		some:     true,
		assigned: true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		some:     false,
		assigned: true,
	}
}

func (o *Option[T]) Some(v T) {
	if o.assigned {
		panic("Option already assigned")
	}
	o.value = v
	o.some = true
	o.assigned = true
}

func (o *Option[T]) None() {
	if o.assigned {
		panic("Option already assigned")
	}
	o.some = false
	o.assigned = true
}

func (o *Option[T]) IsSome() bool {
	return o.some
}

func (o *Option[T]) IsSomeAnd(f func(T) bool) bool {
	if !o.some {
		return false
	}
	return f(o.value)
}

func (o *Option[T]) IsNone() bool {
	return !o.some
}

func (o *Option[T]) IsNoneOr(f func(T) bool) bool {
	if !o.some {
		return true
	}
	return f(o.value)
}

func (o *Option[T]) Expect(msg string) T {
	if o.some {
		return o.value
	}
	panic(msg)
}

func (o *Option[T]) Unwrap() T {
	if o.some {
		return o.value
	}
	panic("called Option.Unwrap() on a None value")
}

func (o *Option[T]) UnwrapOr(def T) T {
	if o.some {
		return o.value
	}
	return def
}

func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if o.some {
		return o.value
	}
	return f()
}

func (o *Option[T]) UnwrapOrDefault() T {
	if o.some {
		return o.value
	}
	var zero T
	return zero
}
