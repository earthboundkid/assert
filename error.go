package assert

import (
	"errors"
)

// OK asserts that error is nil and returns value.
// Typical use is like
//
//	v := be.OK(canFail())
func (be Tester) OK[T any](value T, err error) T {
	be.tb.Helper()
	if err != nil {
		be.fatalf("err != nil: %v", err)
		return *new(T)
	}
	return value
}

// ErrorIs asserts that got [errors.Is] target.
func (be Tester) ErrorIs(got, target error) Tester {
	be.tb.Helper()
	if !errors.Is(got, target) {
		be.fatalf("got errors.Is(%v, %v) == false", got, target)
	}
	return be
}

// ErrorAsType asserts that [errors.AsType] can unwrap got as T.
func (be Tester) ErrorAsType[T error](got error) T {
	be.tb.Helper()
	err, ok := errors.AsType[T](got)
	if !ok {
		be.fatalf("got errors.AsType[%T](%v) == false", *new(T), got)
	}
	return err
}
