package option

type Option[T any] struct {
	value T
	exists bool
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func Some[T any](value T) Option[T] {
	return Option[T] { value, true }
}

func (this Option[T]) IsSome() bool {
	return this.exists
}

func (this *Option[T]) Set(value T) {
	this.value = value
	this.exists = true
}

func (this Option[T]) Get() (T, bool) {
	return this.value, this.exists
}

func (this Option[T]) Unwrap() T {
	val, exists := this.Get()
	if !exists {
		panic("option: Unwrap() called on a None option.")
	}

	return val
}

func (this Option[T]) UnwrapOr(defaultValue T) T {
	val, exists := this.Get()
	if exists {
		return val
	} else {
		return defaultValue
	}
}
