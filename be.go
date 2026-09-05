package assert

import (
	"slices"
	"testing"
)

// Tester is a type that wraps a [*testing.T], [*testing.B], or [*testing.F]
// and adds methods for doing assertion tests against
// that test manager.
//
// The methods of testing.TB that Tester uses are Helper, Logf, and FailNow (for [FailNow]) or Fail (for [Continue]).
type Tester struct {
	tb      testing.TB
	relaxed bool
}

// FailNow returns a Tester will end the test after an assertion failure with [*testing.T.FailNow].
func FailNow(t testing.TB) Tester {
	return Tester{t, false}
}

// Continue returns a Tester will continue testing even after an assertion failure.
// It calls [*testing.T.Fail].
func Continue(t testing.TB) Tester {
	return Tester{t, true}
}

// TB returns the testing.TB the Tester wraps.
// Mostly useful with [RunAll] to access logging.
func (be Tester) TB() testing.TB {
	return be.tb
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
