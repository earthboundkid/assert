package assert

import (
	"slices"
	"testing"
)

// Tester is a type that wraps a [*testing.T] or [*testing.B]
// and adds methods for doing assertion tests against
// that test manager.
//
// The methods of testing.TB that Tester uses are Helper, Logf, and FailNow (for [Strict]) or Fail (for [Relaxed]).
type Tester struct {
	tb      testing.TB
	relaxed bool
}

// Strict returns a Tester will end the test after an assertion failure with t.FailNow().
func Strict(t testing.TB) Tester {
	return Tester{t, false}
}

// Relaxed returns a Tester will continue testing even after an assertion failure.
func Relaxed(t testing.TB) Tester {
	return Tester{t, true}
}

func (be Tester) fatalf(format string, args ...any) {
	be.tb.Helper()
	be.tb.Logf(format, args...)
	if be.relaxed {
		be.tb.Fail()
	} else {
		be.tb.FailNow()
	}
}

// Equal asserts that got == want.
func (be Tester) Equal[T comparable](got, want T) Tester {
	be.tb.Helper()
	if want != got {
		be.fatalf("want: %v; got: %v", want, got)
	}
	return be
}

// NotEqual asserts that got != want.
func (be Tester) NotEqual[T comparable](got, bad T) Tester {
	be.tb.Helper()
	if got == bad {
		be.fatalf("got: %v", got)
	}
	return be
}

// SlicesEqual asserts that slices.Equal(got, want).
func (be Tester) SlicesEqual[T comparable](got, want []T) Tester {
	be.tb.Helper()
	if !slices.Equal(got, want) {
		be.fatalf("got: %v; want: %v", got, want)
	}
	return be
}

// True asserts that value is true.
func (be Tester) True(value bool) Tester {
	be.tb.Helper()
	if !value {
		be.fatalf("got: false")
	}
	return be
}

// False asserts that value is false.
func (be Tester) False(value bool) Tester {
	be.tb.Helper()
	if value {
		be.fatalf("got: true")
	}
	return be
}
