package assert

import (
	"errors"
)

// ErrorIs calls t.Fatalf if got is not want according to [errors.Is].
func (be Tester) ErrorIs(got, want error) Tester {
	be.tb.Helper()
	if !errors.Is(got, want) {
		be.fatalf("got errors.Is(%v, %v) == false", got, want)
	}
	return be
}

// ErrorAs calls t.Fatalf if got cannot be assigned to want by [errors.As].
func (be Tester) ErrorAsType[T error](got error) error {
	be.tb.Helper()
	err, ok := errors.AsType[T](got)
	if !ok {
		be.fatalf("got errors.AsType[%T](%v) == false", *new(T), got)
	}
	return err
}
