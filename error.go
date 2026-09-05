package assert

import (
	"errors"
)

// OK asserts that error is nil and returns value.
// Typical use is like
//
//	v := be.OK(canFail())
func (be TB) OK[T any](value T, err error) T {
	be.Helper()
	if err != nil {
		be.fatalf("err != nil: %v", err)
		return *new(T)
	}
	return value
}

// OK asserts that error is nil and returns v1 and v2.
// Typical use is like
//
//	v1, v2 := be.OK2(canFail())
func (be TB) OK2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	be.Helper()
	if err != nil {
		be.fatalf("err != nil: %v", err)
		return *new(T1), *new(T2)
	}
	return v1, v2
}

// ErrorIs asserts that got [errors.Is] target.
func (be TB) ErrorIs(got, target error) TB {
	be.Helper()
	if !errors.Is(got, target) {
		be.fatalf("got errors.Is(%v, %v) == false", got, target)
	}
	return be
}

// ErrorAsType asserts that [errors.AsType] can unwrap got as T.
func (be TB) ErrorAsType[T error](got error) T {
	be.Helper()
	err, ok := errors.AsType[T](got)
	if !ok {
		be.fatalf("got errors.AsType[%T](%v) == false", *new(T), got)
	}
	return err
}
