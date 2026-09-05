package assert

import (
	"slices"
	"testing"
)

// TB is a type that wraps a [*testing.T], [*testing.B], or [*testing.F]
// and adds methods for doing assertion tests against
// that test manager.
//
// The methods of testing.TB that TB uses for assertions are
// Helper, Logf, and FailNow (for [FailsNow]) or Fail (for [Continues]).
type TB struct {
	testing.TB
	relaxed bool
}

var _ testing.TB = TB{}

// FailsNow returns a TB
// that will end the test after an assertion failure with [testing.TB.FailNow].
func FailsNow(t testing.TB) TB {
	return TB{t, false}
}

// FailsNow returns a copy of the TB
// that will end the test after an assertion failure with [testing.TB.FailNow].
func (be TB) FailsNow() TB {
	return TB{be.TB, false}
}

// Continues returns a TB
// that will continue testing even after an assertion failure.
// It calls [testing.TB.Fail].
func Continues(t testing.TB) TB {
	return TB{t, true}
}

// Continues returns a copy of the TB
// that will continue testing even after an assertion failure.
// It calls [testing.TB.Fail].
func (be TB) Continues() TB {
	return TB{be.TB, true}
}

func (be TB) fatalf(format string, args ...any) {
	be.Helper()
	be.TB.Logf(format, args...)
	if be.relaxed {
		be.TB.Fail()
	} else {
		be.TB.FailNow()
	}
}

// Equal asserts that got == want.
func (be TB) Equal[T comparable](got, want T) TB {
	be.Helper()
	if want != got {
		be.fatalf("want: %v; got: %v", want, got)
	}
	return be
}

// NotEqual asserts that got != want.
func (be TB) NotEqual[T comparable](got, bad T) TB {
	be.Helper()
	if got == bad {
		be.fatalf("got: %v", got)
	}
	return be
}

// SlicesEqual asserts that slices.Equal(got, want).
func (be TB) SlicesEqual[T comparable](got, want []T) TB {
	be.Helper()
	if !slices.Equal(got, want) {
		be.fatalf("got: %v; want: %v", got, want)
	}
	return be
}

// True asserts that value is true.
func (be TB) True(value bool) TB {
	be.Helper()
	if !value {
		be.fatalf("got: false")
	}
	return be
}

// False asserts that value is false.
func (be TB) False(value bool) TB {
	be.Helper()
	if value {
		be.fatalf("got: true")
	}
	return be
}
